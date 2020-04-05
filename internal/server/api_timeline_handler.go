package server

import "github.com/gin-gonic/gin"

func (server *Server) TimeLine() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.IndentedJSON(200, server.JsonData.TimeLineData.TimeLine)
	}
}