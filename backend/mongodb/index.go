package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndex(client *mongo.Client, dbName string) {
	profileCollection := client.Database(dbName).Collection("profiles")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	cursor, err := profileCollection.Indexes().List(context.Background())
	if err != nil {
		log.Fatal("Unable to list indexes for profiles collection : ", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			log.Fatal("Unable to decode index for profiles collection : ", err)
		}
		if index["name"] == "email_1" {
			log.Println("Index already exists for email field in profiles collection")
			return
		}
	}

	_, err = profileCollection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal("Unable to create index for email field in profiles collection : ", err)
	}

	log.Println("Index created for email field in profiles collection")
}
