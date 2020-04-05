package server
import (
	"sort"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	"github.com/gin-gonic/gin"
)

type Locations struct{
	Loc []string `form:"loc"`
}

func (server *Server) Location() gin.HandlerFunc {
	return func(c *gin.Context) {
		var l Locations
		c.Bind(&l)
		if len(l.Loc) > 0 {
			d := make(map[string]interface{})
			for _,v := range l.Loc{
				x := server.JsonData[v]
				if x != nil {
					x = x.(map[string]interface{})
					d[v] = x
				}
			}
			c.JSON(200,d)
		} else {
			c.JSON(200,server.JsonData.Districts.Loc)
		}
	}
}