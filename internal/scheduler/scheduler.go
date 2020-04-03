package scheduler

import (
	"os/exec"

	. "github.com/yedhink/covid19-kerala-api/internal/controller"
	. "github.com/yedhink/covid19-kerala-api/internal/logger"
	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
	. "github.com/yedhink/covid19-kerala-api/internal/server"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
	. "github.com/yedhink/covid19-kerala-api/internal/website"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	CronSpec string
	Sc Scraper
	Site *Website
	St *Storage
	Server *Server
}

func (s Scheduler) Schedule(){
	c := cron.New()
	id,err := c.AddFunc(s.CronSpec, s.BackgroundDaemon)
	if err != nil{
		Log.Printf(Error("cron error : scheduling background daemon failed %v\n",err))
		return
	} else {
		Log.Printf(Success("cron scheduled to run with spec '%s' and id %v\n",s.CronSpec,id))
	}
	c.Run()
	select {}
}

func (s Scheduler) execScript(program string, script string) {
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

func (s Scheduler) BackgroundDaemon(){
	file := s.St.LocalPDFName()
	Log.Printf(Info("Requesting data from dhs kerala website...."))
	res := s.Sc.GetMainPage()
	s.Site.BulletinPageURL = res[1]
	if file == s.St.BasePath+res[0] {
		Log.Printf(Info("The pdf file is already latest"))
	} else {
		s.St.RemoteFileName = res[0]
		Log.Printf(Info("You need latest pdf file : %s(local) != %s(remote)\n", file, s.St.BasePath+res[0]))
		Log.Printf(Info("lastest file : %s\n",s.Sc.GetLatestPDF()))
		err := s.Site.Download(s.St)
		if err != nil {
			Log.Printf(Error("Download Failed! %v",err))
		} else {
			Log.Printf(Success("Downloaded latest pdf file into %s",s.St.BasePath+s.St.RemoteFileName))
			s.execScript("python3", "scripts/extract-text-data.py")
		}
	}
}