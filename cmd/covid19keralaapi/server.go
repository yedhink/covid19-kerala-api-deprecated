package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anaskhan96/soup"
)

type Scraper interface {
	GetMainPage(chan [2]string)
	GetLatestPDF(chan string)
}

type Website struct {
	baseURL          string
	mainPageURL      string
	bulletinPageURL  string
	latestPDFFileURL string
}

func Scrape(base string, route string) soup.Root {
	resp, err := soup.Get(base + route)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return soup.HTMLParse(resp)
}

func (w *Website) GetLatestPDF(c chan string) {
	doc := Scrape(w.baseURL, w.bulletinPageURL)
	fmt.Printf("bullet url : %s\n",w.bulletinPageURL)
	links := doc.Find("div", "class", "entry-content").Find("a")
	w.latestPDFFileURL = links.Attrs()["href"]
	c <- w.baseURL + w.latestPDFFileURL
}

func (w *Website) GetMainPage(c chan [2]string) {
	doc := Scrape(w.baseURL, w.mainPageURL)
	links := doc.Find("h3", "class", "entry-title").Find("a")
	latestFileName := "data/" + strings.ReplaceAll(links.Text(), "/", "-") + ".pdf"
	bulletinPage := links.Attrs()["href"]
	c <- [2]string{latestFileName, bulletinPage}
}

func main() {
	var website = &Website{
		baseURL: "http://dhs.kerala.gov.in/",
		mainPageURL: `%e0%b4%a1%e0%b5%86%e0%b4%af%e0%b4%bf%e0%b4%b2%e0%b4%bf-` +
			`%e0%b4%ac%e0%b5%81%e0%b4%b3%e0%b5%8d%e0%b4%b3%e0%b4%b1` +
			`%e0%b5%8d%e0%b4%b1%e0%b4%bf%e0%b4%a8%e0%b5%8d%e2%80%8d/`,
	}
	var sc Scraper = website

	/*
		Glob ignores file system errors such as I/O errors reading directories.
		The only possible returned error is ErrBadPattern, when pattern is malformed.
	*/
	files, err := filepath.Glob("data/*.pdf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	chan1 := make(chan [2]string)
	chan2 := make(chan string)
	for _, f := range files {
		go sc.GetMainPage(chan1)
		x := <-chan1
		website.bulletinPageURL = x[1]
		if f == x[0] {
			fmt.Println("The pdf file is already latest")
		} else {
			fmt.Printf("You need latest pdf file : %s(local) != %s(remote)\n", f, x[0])
			go sc.GetLatestPDF(chan2)
			fmt.Printf("lastest file : %s\n",<-chan2)
		}
	}
}
