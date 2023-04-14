package api

import (
	"log"

	"github.com/achintya-7/sigma-chat/mongodb"
	"github.com/achintya-7/sigma-chat/sockets"
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
	socketServer := sockets.Setup()

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	defer func() {
		server.router = r
	}()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hi Human!": "Welcome to Sigma Chat",
		})
	})

	auth := r.Group("/auth")
	auth.POST("/signup", server.signUp)
	auth.POST("/login", server.login)

	profile := r.Group("/profile")
	profile.GET("/all/:username")
	profile.GET("/:email")

	chats := r.Group("/chats")
	chats.GET("/:profile", server.getChats)
	chats.POST("/", server.createNewChat)

	messages := r.Group("/messages")
	messages.GET("/:chatId", server.getMessages)
	// we dont really need it as we can add a callback to the socketio callback
	messages.POST("/", server.sendMessage)

	r.GET("/socket.io/*any", gin.WrapH(socketServer))
	r.POST("/socket.io/*any", gin.WrapH(socketServer))
}

func (server *Server) Start(serverAddress string) error {
	log.Printf("Server starting ")
	return server.router.Run(serverAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
