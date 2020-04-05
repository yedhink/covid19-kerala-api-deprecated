package controller

import (
	"io/ioutil"

	"github.com/json-iterator/go"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	. "github.com/yedhink/covid19-kerala-api/internal/model"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

func genarateTimeline(st *Storage,d *Data,t *TimeLine) map[string]interface{}{
	// messily generates a timeline and stores into a map
	// requires refactoring into struct model
	date,_ := time.Parse("01-02-2006", GetLocalPdfDate(st.BasePath))
	latest := fmt.Sprintf("%02d-%02d-%02dT00:00:00Z", date.Year(),date.Month(),date.Day())
	latest_value := d.Data[latest].(map[string]interface{})["total"].(map[string]interface{})["no_of_positive_cases_admitted"]
	t.TimeLine["total_no_of_positive_cases_admitted"] = make(map[string]interface{},2)
	t.TimeLine["total_no_of_positive_cases_admitted"].(map[string]interface{})["latest"] = latest_value
	t.TimeLine["total_no_of_positive_cases_admitted"].(map[string]interface{})["timeline"] = make(map[string]interface{},len(d.Data))
	for k,v := range d.Data {
		t.TimeLine["total_no_of_positive_cases_admitted"].(map[string]interface{})["timeline"].(map[string]interface{})[k] = v.(map[string]interface{})["total"].(map[string]interface{})["no_of_positive_cases_admitted"]
	}
	return t.TimeLine
}

func Deserialize(st *Storage) map[string]interface{}{
	var d Data
	// Open our jsonFile
	file, err := ioutil.ReadFile(st.BasePath+st.JsonFileName)
	if err != nil {
		Log.Printf(Error("failed to read json file : %s\n",err))
	}
	Log.Printf(Success("Successfully Opened %s\n",st.BasePath+st.JsonFileName))
	genarateTimeline(st,&d,&t)
	dataset.TimeData = t
}