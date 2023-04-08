package api

type Server struct {
	config     utils.Config
	router     *gin.Engine
	client     *mongo.Client
	collection utils.Collection
}
