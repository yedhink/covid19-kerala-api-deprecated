package scraper

import (
	"fmt"
	"os"
	"github.com/anaskhan96/soup"
)

type Scraper interface {
	GetMainPage() [2]string
	GetLatestPDF() string
}

func Scrape(base string, route string,attr string,className string) soup.Root {
	resp, err := soup.Get(base + route)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	doc :=  soup.HTMLParse(resp)
	return doc.Find(attr, "class", className).Find("a")
}
