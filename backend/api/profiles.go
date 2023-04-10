package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type searchAll struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) searchProfile(c *gin.Context) {
	var req searchAll
	var resp []SignUp

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	regex := bson.M{
		"$regex":  "^" + req.Username,
		"options": "i",
	}

	filter := bson.M{"username": regex}

	opts := options.Find().SetLimit(20)

	cursor, err := server.collection.Profiles.Find(c, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = cursor.All(c, &resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(errors.New("Unable to parsse resp")))
		return
	}

	c.JSON(http.StatusOK, resp)
}

type search struct {
	Email string `uri:"email" binding:"required"`
}

type profileResp struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Email    string             `json:"email" binding:"required"`
	Bio      string             `json:"bio" binding:"required"`
}

func (server *Server) getOneProfile(c *gin.Context) {
	var req search
	var resp profileResp

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	filter := bson.M{"email": req.Email}

	err := server.collection.Profiles.FindOne(c, filter).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, errorResponse(errors.New("No Profile found")))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, resp)
}

