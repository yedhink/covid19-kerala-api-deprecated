package server

import (
	"github.com/gin-gonic/gin"
	. "github.com/yedhink/covid19-kerala-api/internal/controller"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

type Server struct {
	Port string
	JsonData map[string]interface{}
}

func (server *Server) Start(st *Storage) {
	server.JsonData = Deserialize(st)
	router := gin.Default()
	router.GET("/api", server.Api())
	router.GET("/api/location", server.Location())
	router.Run() // listen and serve on 0.0.0.0:8080
}
