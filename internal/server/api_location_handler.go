package server
import (
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
	"time"
	"github.com/gin-gonic/gin"
)

type Locations struct{
	Loc []string `form:"loc"`
	Date string `form:"date"`
}

func filterByLoc(value map[string]interface{},key string,d map[string]interface{},userLoc []string) {
	for _,loc := range userLoc {
		d[key].(map[string]interface{})[loc] = value[loc]
	}
}

func parseDate(k string, s string) (time.Time,time.Time) {
	userDate,err := time.Parse("02-01-2006", s)
	if err != nil {
		userDate,_ =  time.Parse("02/01/2006", s)
	}
	userDate.Format(time.RFC3339)
	keyDate,_ := time.Parse(time.RFC3339,k)
	return keyDate,userDate
}

func validateDate(k string,d string,st *Storage) bool {
	switch d[:1] {
	case "<":
		kD,uD := parseDate(k,d[1:])
		return kD.Before(uD)
	case ">":
		kD,uD := parseDate(k,d[1:])
		return kD.After(uD)
	case "l":
		// latest data pdf
		date := GetLocalPdfDate(st.BasePath)
		kD,uD := parseDate(k, date)
		return kD.Equal(uD)
	default:
		kD,uD := parseDate(k,d)
		return kD.Equal(uD)
	}
}


func (server *Server) Location(st *Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var l Locations
		c.Bind(&l)
		if len(l.Loc) > 0 {
			d := make(map[string]interface{})
			for key,value := range server.JsonData.All.Data{
				if l.Date != "" {
					if validateDate(key,l.Date,st) {
						d[key] = make(map[string]interface{},len(l.Loc))
						filterByLoc(value.(map[string]interface{}),key,d,l.Loc)
					}
				} else {
					d[key] = make(map[string]interface{},len(l.Loc))
					filterByLoc(value.(map[string]interface{}),key,d,l.Loc)
				}
			}
			c.IndentedJSON(200,d)
		} else {
			if l.Date != "" {
				d := make(map[string]interface{})
				for key,value := range server.JsonData.All.Data{
					if validateDate(key,l.Date,st) {
						d[key] = value
					}
				}
				c.IndentedJSON(200,d)
			} else {
				c.IndentedJSON(200,server.JsonData.Districts.Loc)
			}
		}
	}
}