package main

import (
	"log"

	"github.com/achintya-7/sigma-chat/api"
	"github.com/achintya-7/sigma-chat/mongodb"
	"github.com/achintya-7/sigma-chat/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	mongoClient, err := mongodb.ConnectDB(config.MongoUri)
	if err != nil {
		log.Fatal("error connecting to DB: ", err)
	}

	server := api.NewServer(config, mongoClient)
	if err != nil {
		log.Fatal("error creating server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
