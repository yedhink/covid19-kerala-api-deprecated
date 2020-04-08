package server

import "github.com/gin-gonic/gin"

func (server *Server) NoRouteErr() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(404, gin.H{"success": false, "message": "Page not found"})
	}
}
