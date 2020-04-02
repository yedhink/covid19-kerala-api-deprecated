package website

import (
	"strings"
	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
)

type Website struct {
	BaseURL          string
	MainPageURL      string
	BulletinPageURL  string
	LatestPDFURL string
}

func (w *Website) GetLatestPDF() (string){
	links := Scrape(w.BaseURL, w.BulletinPageURL,"div","entry-content")
	w.LatestPDFURL = links.Attrs()["href"]
	return w.BaseURL + w.LatestPDFURL
}

func (w *Website) GetMainPage() ([2]string){
	links := Scrape(w.BaseURL, w.MainPageURL,"h3","entry-title")
	latestFileName := strings.ReplaceAll(links.Text(), "/", "-") + ".pdf"
	bulletinPage := links.Attrs()["href"]
	return [2]string{latestFileName, bulletinPage}
}
