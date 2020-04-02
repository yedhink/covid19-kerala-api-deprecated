package main

import (
	. "github.com/yedhink/covid19-kerala-api/internal/scheduler"
	"github.com/robfig/cron/v3"
)


func main() {
	c := cron.New()
	c.AddFunc("* 15-16 * * *", BackgroundDaemon)
	c.Run()
	select {}
}
