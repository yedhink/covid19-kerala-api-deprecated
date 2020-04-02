package scheduler

import (
	"fmt"
	"os/exec"

	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
	. "github.com/yedhink/covid19-kerala-api/internal/website"
)

var website = &Website{
	BaseURL: "http://dhs.kerala.gov.in",
	MainPageURL: `/%e0%b4%a1%e0%b5%86%e0%b4%af%e0%b4%bf%e0%b4%b2%e0%b4%bf-` +
		`%e0%b4%ac%e0%b5%81%e0%b4%b3%e0%b5%8d%e0%b4%b3%e0%b4%b1` +
		`%e0%b5%8d%e0%b4%b1%e0%b4%bf%e0%b4%a8%e0%b5%8d%e2%80%8d/`,
}
var sc Scraper = website
var st = &Storage{
	BasePath : "data/",
}

func execScript(program string, script string) {
	// hardcoded since the scripts dir under pipenv
	cmd := exec.Command(program, script,"-w")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out)
}

func BackgroundDaemon(){
	file := st.LocalPDFName()
	res := sc.GetMainPage()
	website.BulletinPageURL = res[1]
	if file == st.BasePath+res[0] {
		fmt.Println("The pdf file is already latest")
	} else {
		st.RemoteFileName = res[0]
		fmt.Printf("You need latest pdf file : %s(local) != %s(remote)\n", file, st.BasePath+res[0])
		fmt.Printf("lastest file : %s\n",sc.GetLatestPDF())
		err := website.Download(st)
		if err != nil {
			fmt.Println(err)
		} else {
			execScript("python3", "scripts/extract-text-data.py")
		}
	}
}