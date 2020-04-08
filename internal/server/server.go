package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	. "github.com/yedhink/covid19-kerala-api/internal/model"
	. "github.com/yedhink/covid19-kerala-api/internal/controller"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

type Server struct {
	Port string
	JsonData DataSet
}

func (server *Server) Start(st *Storage) {
	server.JsonData = Deserialize(st)
	if server.Port == "" {
		Log.Printf(Error("PORT env variable must be set in shell before executing the binary : eg:- PORT=5000 ./main"))
		return
	}

	// run in Release mode by default
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(),gin.Recovery())
	router.Use(favicon.New("web/assets/favicon.ico"))
	router.LoadHTMLFiles("web/index.html")
	router.GET("/", server.Root())
	router.GET("/api", server.Api())
	router.GET("/api/location", server.Location(st))
	router.GET("/api/timeline", server.TimeLine())
	router.NoRoute(server.NoRouteErr())
	router.Run(server.Port) // listen and serve on 0.0.0.0:8080 by default
}
