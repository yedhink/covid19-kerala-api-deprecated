package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
	JsonData map[string]interface{}
}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}