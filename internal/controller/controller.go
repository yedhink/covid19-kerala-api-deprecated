package controller

import (
	"io/ioutil"
	"fmt"
	"time"
	"sort"

	"github.com/json-iterator/go"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	. "github.com/yedhink/covid19-kerala-api/internal/model"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

func getLocations(v map[string]interface{}) map[string]interface{} {
	d := make(map[string]interface{})
	d["locations"] = make([]string,0)
	for k,_ := range v {
		d["locations"] = append(d["locations"].([]string),k)
	}
	sort.Strings(d["locations"].([]string))
	return d
}

func genarateTimeline(st *Storage,d *Data,t *TimeLine,l *Location) map[string]interface{} {
	// messily generates a timeline and stores into a map
	// requires refactoring into struct model

	//                    mm-dd-yyyy
	date,_ := time.Parse("02-01-2006", GetLocalPdfDate(st.BasePath))
	latest_date := fmt.Sprintf("%02d-%02d-%02dT00:00:00Z", date.Year(),date.Month(),date.Day())
	latest_obj := d.Data[latest_date].(map[string]interface{})
	latest_value := latest_obj["total"].(map[string]interface{})["no_of_positive_cases_admitted"]

	t.TimeLine["total_no_of_positive_cases_admitted"] = make(map[string]interface{},2)
	t.TimeLine["total_no_of_positive_cases_admitted"].(map[string]interface{})["latest"] = latest_value

	t.TimeLine["total_no_of_positive_cases_admitted"].(map[string]interface{})["timeline"] = make(map[string]interface{},len(d.Data))

	for k,v := range d.Data {
		t.TimeLine["total_no_of_positive_cases_admitted"].(map[string]interface{})["timeline"].(map[string]interface{})[k] = v.(map[string]interface{})["total"].(map[string]interface{})["no_of_positive_cases_admitted"]
	}
	l.Loc = getLocations(latest_obj)
	return t.TimeLine
}

func Deserialize(st *Storage) DataSet{
	var dataset DataSet
	var d = NewData()
	var t = NewTimeLine()
	var l = NewLocation()
	// Open our jsonFile
	file, err := ioutil.ReadFile(st.BasePath+st.JsonFileName)
	if err != nil {
		Log.Printf(Error("failed to read json file : %s\n",err))
	}
	Log.Printf(Success("Successfully Opened %s\n",st.BasePath+st.JsonFileName))
	// read the whole json
	err = jsoniter.Unmarshal(file, &d.Data)
	if err != nil {
		Log.Printf(Error("failed to unmarshal into json obj : %s\n",err))
	}
	dataset.All = d
	// generate the timeline into a map
	genarateTimeline(st,&d,&t,&l)
	dataset.TimeLineData = t
	dataset.Districts = l
	return dataset
}