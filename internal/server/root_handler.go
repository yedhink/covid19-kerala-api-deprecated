package server

import "github.com/gin-gonic/gin"

func (server *Server) Root() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	}
}