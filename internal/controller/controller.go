package controller

import (
	"io/ioutil"

	"github.com/json-iterator/go"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	. "github.com/yedhink/covid19-kerala-api/internal/model"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)


func Deserialize(st *Storage) map[string]interface{}{
	var d Data
	// Open our jsonFile
	file, err := ioutil.ReadFile(st.BasePath+st.JsonFileName)
	if err != nil {
		Log.Printf(Error("failed to read json file : %s\n",err))
	}
	Log.Printf(Success("Successfully Opened %s\n",st.BasePath+st.JsonFileName))
	jsoniter.Unmarshal(file, &d.Data)
	return d.Data
}