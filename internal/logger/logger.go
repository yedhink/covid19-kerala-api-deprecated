package logger

import (
	"io"
	"log"
	"os"
	"github.com/fatih/color"
)

var (
	Log *log.Logger
	err = color.New(color.FgRed)
	info = color.New(color.FgGreen).Add(color.BgHiYellow)
)

func Error(format string,v ...interface{}) string{
	return err.Sprintf(format, v...)
}

func Info(format string,v ...interface{}) string{
	return info.Sprintf(format, v...)
}

func init() {
	fileLog, err := os.OpenFile("/tmp/covid19-kerala-api.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	mw1 := io.MultiWriter(os.Stdout, fileLog)
	Log = log.New(mw1, "scheduler running :", log.Ldate|log.Ltime|log.Lshortfile)
}