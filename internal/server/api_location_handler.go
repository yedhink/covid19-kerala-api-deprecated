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
		Log.Print(Error("len of loc : %d",len(l.Loc)))
		if len(l.Loc) > 0 {
			x := server.JsonData[l.Loc[0]]
			if x != nil {
				x = x.(map[string]interface{})
			}
			c.JSON(200,x)
		} else {
			d := make(map[string][]string)
			for k,_ := range server.JsonData {
				if k != "total" {
					d["locations"] = append(d["locations"],k)
				}
			}
			sort.Strings(d["locations"])
			c.JSON(200,d)
		}
	}
}