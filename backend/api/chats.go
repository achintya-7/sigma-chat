package api

import (
	"errors"
	"net/http"

	"github.com/achintya-7/sigma-chat/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type createChat struct {
	profile1     string `json:"profile1" binding:"required"`
	prfoile1Name string `json:"profile1name" binding:"required"`
	profile2     string `json:"profile2" binding:"required"`
	profile2Name string `json:"profile2name" binding:"required"`
}

func (server *Server) createNewChat(c *gin.Context) {
	var resp models.Chat
	var req createChat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.profile1 == req.profile2 {
		c.JSON(http.StatusBadRequest, errorResponse(errors.New("profileId1 and profileId2 are same")))
		return
	}

	if req.profile1 > req.profile2 {
		resp.ChatId = req.profile1 + req.profile2
	} else {
		resp.ChatId = req.profile2 + req.profile1
	}

	resp.Profile1 = req.profile1
	resp.Profile1Name = req.prfoile1Name
	resp.Profile2 = req.profile2
	resp.Profile2Name = req.profile2Name

	_, err := server.collection.Chats.InsertOne(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

type getChat struct {
	profile string `uri:"profile" binding:"required"`
}

func (server *Server) getChat(c *gin.Context) {
	var resp []models.Chat
	var req getChat
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
		c.JSON(http.StatusInternalServerError, errorResponse(errors.New("Unable to parse resp")))
		return
	}

	c.JSON(http.StatusOK, resp)
}
