package scheduler

import (
	"os/exec"

	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
	. "github.com/yedhink/covid19-kerala-api/internal/website"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	Spec string
}

var website = &Website{
	BaseURL: "http://dhs.kerala.gov.in",
	MainPageURL: `/%e0%b4%a1%e0%b5%86%e0%b4%af%e0%b4%bf%e0%b4%b2%e0%b4%bf-` +
		`%e0%b4%ac%e0%b5%81%e0%b4%b3%e0%b5%8d%e0%b4%b3%e0%b4%b1` +
		`%e0%b5%8d%e0%b4%b1%e0%b4%bf%e0%b4%a8%e0%b5%8d%e2%80%8d/`,
}
var sc Scraper = website
var st = &Storage{
	BasePath : "data/",
	LocalFileExist: false,
}

func (s Scheduler) Schedule(){
	c := cron.New()
	id,err := c.AddFunc(s.Spec, BackgroundDaemon)
	if err != nil{
		Log.Print(Error("cron error : scheduling background daemon failed %v\n",err))
		return
	} else {
		Log.Print(Info("cron scheduled to run with spec %s and id %v\n",s.Spec,id))
	}
	c.Run()
	select {}
}

func execScript(program string, script string) {
	// hardcoded since the scripts dir under pipenv
	cmd := exec.Command(program, script,"-w")
	out, err := cmd.CombinedOutput()
	if err != nil {
		Log.Print(Error("Failed to execute exec python script",err))
	}
	Log.Print(Info("%s", out))
}

func BackgroundDaemon(){
	file := st.LocalPDFName()
	Log.Print(Info("Requesting data from dhs kerala website...."))
	res := sc.GetMainPage()
	website.BulletinPageURL = res[1]
	if file == st.BasePath+res[0] {
		Log.Print(Info("The pdf file is already latest"))
	} else {
		st.RemoteFileName = res[0]
		Log.Print(Info("You need latest pdf file : %s(local) != %s(remote)\n", file, st.BasePath+res[0]))
		Log.Print(Info("lastest file : %s\n",sc.GetLatestPDF()))
		err := website.Download(st)
		if err != nil {
			Log.Print(Error("Download Failed!",err))
		} else {
			execScript("python3", "scripts/extract-text-data.py")
		}
	}
}