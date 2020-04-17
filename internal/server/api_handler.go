package server
import "github.com/gin-gonic/gin"

func (server *Server) Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		server.JsonData.All.Data["success"] = true
		c.JSON(200, server.JsonData.All.Data)
	}
}
