package mongodb

import (
	"github.com/achintya-7/sigma-chat/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	Chats    *mongo.Collection
	Profiles *mongo.Collection
}

func NewCollection(client *mongo.Client, config utils.Config) Collection {
	return Collection{
		Chats:    client.Database(config.DBName).Collection("chats"),
		Profiles: client.Database(config.DBName).Collection("profiles"),
	}
}
