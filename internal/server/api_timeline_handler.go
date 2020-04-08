package server

import "github.com/gin-gonic/gin"

func (server *Server) TimeLine() gin.HandlerFunc{
	return func(c *gin.Context) {
		server.JsonData.TimeLineData.TimeLine["success"]= true
		c.IndentedJSON(200, server.JsonData.TimeLineData.TimeLine)
	}
}