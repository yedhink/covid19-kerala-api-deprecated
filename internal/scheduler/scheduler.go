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
	CronSpec string
	Sc Scraper
	Site *Website
	St *Storage
	Server *Server
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
	JsonFileName: "data.json",
	LocalFileExist: false,
}

func (s Scheduler) Schedule(){
	c := cron.New()
	id,err := c.AddFunc(s.Spec, BackgroundDaemon)
	if err != nil{
		Log.Printf(Error("cron error : scheduling background daemon failed %v\n",err))
		return
	} else {
		Log.Printf(Success("cron scheduled to run with spec '%s' and id %v\n",s.Spec,id))
	}
	c.Run()
	select {}
}

func execScript(program string, script string) {
	// the program and script locations are hardcoded for now
	// not platform agnostic at the moment
	cmd := exec.Command(program, script,"-w")
	out, err := cmd.CombinedOutput()
	if err != nil {
		Log.Printf(Error("Failed to execute exec python script %v",err))
		Log.Printf(Error("python script output : %s", out))
		Log.Printf(Error("Make sure you've enabled pipenv shell in 'scripts' folder"))
	} else {
		s.Server.JsonData = Deserialize(s.St)
		Log.Printf(Success("%s", out))
	}
}

func BackgroundDaemon(){
	file := st.LocalPDFName()
	Log.Printf(Info("Requesting data from dhs kerala website...."))
	res := sc.GetMainPage()
	website.BulletinPageURL = res[1]
	if file == st.BasePath+res[0] {
		Log.Printf(Info("The pdf file is already latest"))
	} else {
		st.RemoteFileName = res[0]
		Log.Printf(Info("You need latest pdf file : %s(local) != %s(remote)\n", file, st.BasePath+res[0]))
		Log.Printf(Info("lastest file : %s\n",sc.GetLatestPDF()))
		err := website.Download(st)
		if err != nil {
			Log.Printf(Error("Download Failed! %v",err))
		} else {
			Log.Printf(Success("Downloaded latest pdf file into %s",st.BasePath+st.RemoteFileName))
			execScript("python3", "scripts/extract-text-data.py")
		}
	}
}