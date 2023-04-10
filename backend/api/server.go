package api

import (
	"github.com/achintya-7/sigma-chat/mongodb"
	"github.com/achintya-7/sigma-chat/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config     utils.Config
	router     *gin.Engine
	client     *mongo.Client
	collection mongodb.Collection
}

func NewServer(config utils.Config, client *mongo.Client) *Server {

	collection := mongodb.NewCollection(client, config)

	server := &Server{
		config: config,
		client: client,
		collection: collection,
	}

	setupRoutes(server)

	return server
}

func setupRoutes(server *Server) {
	r := gin.Default()

	defer func()  {
		server.router = r
	}()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hi Human!": "Welcome to Sigma Chat",
		})
	})
}

func (server *Server) Start(serverAddress string) error {
	return server.router.Run(serverAddress)
}