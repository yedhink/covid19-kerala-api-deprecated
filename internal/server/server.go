package server

import (
	. "github.com/yedhink/covid19-kerala-api/internal/controller"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
	JsonData map[string]interface{}
}

func (server *Server) Start(st *Storage) {
	server.JsonData = Deserialize(st)
	router := gin.Default()
	router.GET("/api", server.Api())
	router.GET("/api/location", server.Location())
	router.Run() // listen and serve on 0.0.0.0:8080
}