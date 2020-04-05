package server
import (
	"sort"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	"github.com/gin-gonic/gin"
)

type Locations struct{
	Loc []string `form:"loc"`

func filterByLoc(value map[string]interface{},key string,d map[string]interface{},userLoc []string) {
	for _,loc := range userLoc {
		d[key].(map[string]interface{})[loc] = value[loc]
	}
}

}

func (server *Server) Location() gin.HandlerFunc {
	return func(c *gin.Context) {
		var l Locations
		c.Bind(&l)
		if len(l.Loc) > 0 {
			d := make(map[string]interface{})
						filterByLoc(value.(map[string]interface{}),key,d,l.Loc)
						filterByLoc(value.(map[string]interface{}),key,d,l.Loc)
				}
			}
			c.IndentedJSON(200,d)
		} else {
			c.IndentedJSON(200,server.JsonData.Districts.Loc)
		}
	}
}