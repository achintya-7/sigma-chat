package api

import (
	"errors"
	"net/http"

	"github.com/achintya-7/sigma-chat/models"
	"github.com/achintya-7/sigma-chat/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type createChat struct {
	Profile1     string `json:"profile1" binding:"required"`
	Prfoile1Name string `json:"profile1name" binding:"required"`
	Profile2     string `json:"profile2" binding:"required"`
	Profile2Name string `json:"profile2name" binding:"required"`
}

func (server *Server) createNewChat(c *gin.Context) {
	var resp models.Chat
	var req createChat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Profile1 == req.Profile2 {
		c.JSON(http.StatusBadRequest, errorResponse(errors.New("profileId1 and profileId2 are same")))
		return
	}

	if req.Profile1 > req.Profile2 {
		resp.ChatId = req.Profile1 + req.Profile2
	} else {
		resp.ChatId = req.Profile2 + req.Profile1
	}

	resp.Profile1 = req.Profile1
	resp.Profile1Name = req.Prfoile1Name
	resp.Profile2 = req.Profile2
	resp.Profile2Name = req.Profile2Name

	_, err := server.collection.Chats.InsertOne(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

type getChats struct {
	profile string `uri:"profile" binding:"required"`
}

func (server *Server) getChats(c *gin.Context) {
	var resp []models.Chat
	var req getChats
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	filter := bson.M{
		"$or": []bson.M{
			{"profile1": req.profile},
			{"profile2": req.profile},
		},
	}

	cursor, err := server.collection.Chats.Find(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = cursor.All(c, &resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(errors.New("unable to parse resp")))
		return
	}

	c.JSON(http.StatusOK, resp)
}

type getMessages struct {
	chatId string `uri:"chatId" binding:"required"`
}

func (server *Server) getMessages(c *gin.Context) {
	var resp models.Chat
	var req getMessages
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	filter := bson.M{"chatId": req.chatId}

	err := server.collection.Chats.FindOne(c, filter).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, errorResponse(errors.New("chatId not found")))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

}

type SendMessage struct {
	ChatId string `uri:"chatId" binding:"required"`
	Text   string `json:"text" binding:"required"`
	Sender string `json:"sender" binding:"required"`
}

func (server *Server) sendMessage(c *gin.Context) {
	var req SendMessage
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	filter := bson.M{"chatId": req.ChatId}

	update := bson.M{
		"$push": bson.M{
			"messages": bson.M{
				"text":   req.Text,
				"sender": req.Sender,
				"time":   utils.GetTime(),
			},
		},
	}

	resp, err := server.collection.Chats.UpdateMany(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if resp.MatchedCount == 0 {
		c.JSON(http.StatusBadRequest, errorResponse(errors.New("chatId not found")))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "message added",
	})

}
