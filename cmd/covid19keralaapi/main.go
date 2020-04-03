package main

import (
	. "github.com/yedhink/covid19-kerala-api/internal/scheduler"
)


func main() {
	sc := Scheduler{
		Spec : "* 18-20 * * *",
	}
	sc.Schedule()
}
