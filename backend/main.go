package main

import (
	"log"

	"github.com/achintya-7/sigma-chat/api"
	"github.com/achintya-7/sigma-chat/mongodb"
	"github.com/achintya-7/sigma-chat/sockets"
	"github.com/achintya-7/sigma-chat/utils"
)

func main() {
	debug := true

	if debug {
		sockets.Start()
	} else {
		// config, err := utils.LoadConfig(".")
	config, err := utils.OSLoadConfig()
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	mongoClient, err := mongodb.ConnectDB(config.MongoUri, config.DBName)
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

	

}
