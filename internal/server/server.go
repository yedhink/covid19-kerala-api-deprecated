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
	r := gin.Default()
	// currently server.JsonData is not updated dynamically when cron is run
	// maybe move this handler function to outside
	r.GET("/api", func(c *gin.Context) {
		c.IndentedJSON(200, server.JsonData)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}