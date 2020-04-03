package main

import (
	. "github.com/yedhink/covid19-kerala-api/internal/scheduler"
	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
	. "github.com/yedhink/covid19-kerala-api/internal/website"
	. "github.com/yedhink/covid19-kerala-api/internal/server"
)


func main() {
	var website = &Website{
		BaseURL: "http://dhs.kerala.gov.in",
		MainPageURL: `/%e0%b4%a1%e0%b5%86%e0%b4%af%e0%b4%bf%e0%b4%b2%e0%b4%bf-` +
			`%e0%b4%ac%e0%b5%81%e0%b4%b3%e0%b5%8d%e0%b4%b3%e0%b4%b1` +
			`%e0%b5%8d%e0%b4%b1%e0%b4%bf%e0%b4%a8%e0%b5%8d%e2%80%8d/`,
	}

	var scraper Scraper = website

	var storage = &Storage{
		BasePath : "data/",
		JsonFileName: "data.json",
		LocalFileExist: false,
	}

	var server = &Server{
		Port : 8000,
	}

	var scheduler = Scheduler{
		CronSpec : "* * * * *",
		Sc : scraper,
		St : storage,
		Site : website,
		Server : server,
	}


	go scheduler.Schedule()
	server.Start(storage)
}
