package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(mongoUri string, dbName string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Println("cannot create new client", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Println("cannot connect to DB", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("cannot ping to DB", err)
		return nil, err
	}

	createIndex(client, dbName)

	return client, nil
}

func NewMockClient() *mongo.Client {
	return nil
}