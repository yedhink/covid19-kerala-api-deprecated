package main

import (
	. "github.com/yedhink/covid19-kerala-api/internal/scheduler"
	server "github.com/yedhink/covid19-kerala-api/internal/server"
)


func main() {
	sc := Scheduler{
		Spec : "* * * * *",
	}
	go sc.Schedule()
	select {}
	server.Start()
}
