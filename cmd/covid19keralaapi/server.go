package main

import (
	"fmt"
	"strings"

	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

type Website struct {
	baseURL          string
	mainPageURL      string
	bulletinPageURL  string
	latestPDFFileURL string
}

func (w *Website) GetLatestPDF() (string){
	links := Scrape(w.baseURL, w.bulletinPageURL,"div","entry-content")
	w.latestPDFFileURL = links.Attrs()["href"]
	return w.baseURL + w.latestPDFFileURL
}

func (w *Website) GetMainPage() ([2]string){
	links := Scrape(w.baseURL, w.mainPageURL,"h3","entry-title")
	latestFileName := strings.ReplaceAll(links.Text(), "/", "-") + ".pdf"
	bulletinPage := links.Attrs()["href"]
	return [2]string{latestFileName, bulletinPage}
}

func main() {
	var website = &Website{
		baseURL: "http://dhs.kerala.gov.in",
		mainPageURL: `/%e0%b4%a1%e0%b5%86%e0%b4%af%e0%b4%bf%e0%b4%b2%e0%b4%bf-` +
			`%e0%b4%ac%e0%b5%81%e0%b4%b3%e0%b5%8d%e0%b4%b3%e0%b4%b1` +
			`%e0%b5%8d%e0%b4%b1%e0%b4%bf%e0%b4%a8%e0%b5%8d%e2%80%8d/`,
	}
	var sc Scraper = website
	var st = &Storage{
		BasePath : "data/",
	}

	files := st.LocalPDFName()

	for _, f := range files {
		res := sc.GetMainPage()
		website.bulletinPageURL = res[1]
		if f == res[0] {
			fmt.Println("The pdf file is already latest")
		} else {
			st.RemoteFileName = res[0]
			fmt.Printf("You need latest pdf file : %s(local) != %s(remote)\n", f, res[0])
			fmt.Printf("lastest file : %s\n",sc.GetLatestPDF())
			st.Download()
		}
	}
}
