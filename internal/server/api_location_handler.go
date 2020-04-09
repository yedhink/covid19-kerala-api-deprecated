package server

import (
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

type Locations struct {
	Loc  []string `form:"loc"`
	Date string   `form:"date"`
}

func filterByLoc(value map[string]interface{}, key string, d map[string]interface{}, userLoc []string) {
	for _, loc := range userLoc {
		d[key].(map[string]interface{})[loc] = value[loc]
	}
}

func parseDate(k string, s string) (time.Time, time.Time) {
	userDate, err := time.Parse("02-01-2006", s)
	if err != nil {
		userDate, _ = time.Parse("02/01/2006", s)
	}
	userDate.Format(time.RFC3339)
	keyDate, _ := time.Parse(time.RFC3339, k)
	return keyDate, userDate
}
	// if the date is not correctly formatted
	case false:
		d["success"] = false
		d["message"] = "Invalid Date Format as Parameter"
	}
}

func (server *Server) Location(st *Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var l Locations
		c.Bind(&l)
		if len(l.Loc) > 0 {
			d := make(map[string]interface{})
			d["success"] = true
			for key, value := range server.JsonData.All.Data {
				// since we cant assert a value interface of type bool as map[string]interface{}
				if key == "success" {
					continue
				}
				if l.Date != "" {
					if validateDate(key, l.Date, st) {
						d[key] = make(map[string]interface{}, len(l.Loc))
						filterByLoc(value.(map[string]interface{}), key, d, l.Loc)
					}
				} else {
					d[key] = make(map[string]interface{}, len(l.Loc))
					filterByLoc(value.(map[string]interface{}), key, d, l.Loc)
				}
			}
			c.IndentedJSON(200, d)
		} else {
			if l.Date != "" {
				d := make(map[string]interface{})
				d["success"] = true
				for key, value := range server.JsonData.All.Data {
					if key != "success" && validateDate(key, l.Date, st) {
						d[key] = value
					}
				}
				c.IndentedJSON(200, d)
			} else {
				server.JsonData.Districts.Loc["success"] = true
				c.IndentedJSON(200, server.JsonData.Districts.Loc)
			}
		}
	}
}
