package server

import (
	"github.com/gin-gonic/gin"
	. "github.com/yedhink/covid19-kerala-api/internal/date"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

type Query struct {
	Loc  []string `form:"loc"`
	Date string   `form:"date"`
}

// goes through all the input loc and append matches from value obj to the obj d
// value is server.JsonData.All.Data[timestamp], d is the result and userLoc contains all loc params
func filterByLoc(value map[string]interface{}, d map[string]interface{}, userLoc []string) {
	for _, loc := range userLoc {
		d[loc] = value[loc]
	}
}

// q is the Query struct, d is our final result, st is Storage referenc and apiData is server.JsonData.All.Data
// the function validates the date and filters using loc and date parameters of user
func LocDateFilter(q *Query, d map[string]interface{}, st *Storage, apiData map[string]interface{}, locExist bool) {
	// start by checking if the date param value is valid formatted Date
	switch IsDate(q.Date) {
	case true:
		for key, value := range apiData {
			if ValidDate(key, q.Date, st) {
				// if loc length > 0 then only filterByLoc
				if locExist == true {
					d[key] = make(map[string]interface{}, len(q.Loc))
					filterByLoc(value.(map[string]interface{}), d[key].(map[string]interface{}), q.Loc)
				} else {
					// if no loc, then just assign to valid date
					d[key] = value
				}
			}
		}
	// if the date is not correctly formatted
	case false:
		d["success"] = false
		d["message"] = "Invalid Date Format as Parameter"
	}
}

// handles the /location related parameter and non-parametric queries
func (server *Server) Location(st *Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q Query
		c.Bind(&q)
		// initialize the final result 'd' with success=true
		d := make(map[string]interface{})
		d["success"] = true
		// check if loc param is nil or not
		switch length := len(q.Loc); {
		//serve the array of locations
		case length == 0 && q.Date == "":
			d = server.JsonData.Districts.Loc
		// validate just the date and append matching obj entries
		case length == 0 && q.Date != "":
			LocDateFilter(&q, d, st, server.JsonData.All.Data, false)
		// just filter by loc parameter
		case length > 0 && q.Date == "":
			for key, value := range server.JsonData.All.Data {
				d[key] = make(map[string]interface{}, len(q.Loc))
				filterByLoc(value.(map[string]interface{}), d[key].(map[string]interface{}), q.Loc)
			}
		// filter by both loc and date params
		case length > 0 && q.Date != "":
			LocDateFilter(&q, d, st, server.JsonData.All.Data, true)
		}
		c.IndentedJSON(200, d)
	}
}
