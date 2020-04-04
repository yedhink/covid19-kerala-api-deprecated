package server

import (
	"github.com/gin-gonic/gin"
	. "github.com/yedhink/covid19-kerala-api/internal/controller"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

type Server struct {
	Port string
	JsonData map[string]interface{}
}

func (server *Server) Start(st *Storage) {
	server.JsonData = Deserialize(st)
	if server.Port == "" {
		Log.Printf(Error("PORT env variable must be set in shell before executing the binary : eg:- PORT=5000 ./main"))
		return
	}
	router := gin.New()
	router.Use(gin.Logger(),gin.Recovery())
	router.LoadHTMLFiles("web/template/index.html")
	router.GET("/", server.Root())
	router.GET("/api", server.Api())
	router.GET("/api/location", server.Location())
	router.Run(server.Port) // listen and serve on 0.0.0.0:8080
}
