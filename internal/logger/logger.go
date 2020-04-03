package logger

import (
	"io"
	"log"
	"os"
	"github.com/fatih/color"
)

var (
	Log *log.Logger
	err = color.New(color.FgWhite).Add(color.BgRed)
	info = color.New(color.FgHiBlack).Add(color.BgHiYellow)
	success = color.New(color.FgWhite).Add(color.BgGreen)
)

func Error(format string,v ...interface{}) string{
	return err.Sprintf(format, v...)
}

func Info(format string,v ...interface{}) string{
	return info.Sprintf(format, v...)
}

func Success(format string,v ...interface{}) string{
	return success.Sprintf(format, v...)
}


func init() {
	// open the file in readwrite mode | create it if not exist | append to end of file | perm : 0666
	// currently the color ascci code is also appended as its to log file
	// I havent figured out how to disable it on a io.Multiwriter
	fileLog, err := os.OpenFile("/tmp/covid19-kerala-api-scheduler.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	mw1 := io.MultiWriter(os.Stdout, fileLog)
	Log = log.New(mw1, "", log.Ldate|log.Ltime|log.Lshortfile)
}