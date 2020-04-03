package website

import (
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"
	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
	storage "github.com/yedhink/covid19-kerala-api/internal/storage"
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

func (w *Website) Download(st *storage.Storage) error{
	st.Delete()
	resp,err := http.Get(w.BaseURL+w.LatestPDFURL)
	if err != nil {
		return HttpError{err,"http : client policy error or connectivity problems"}
	} else if resp.StatusCode >= 404{
		return HttpError{resp.Status,"http error : failed to GET pdf from "+w.BaseURL+w.LatestPDFURL}
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Trouble reading reesponse body!")
	}
	err = ioutil.WriteFile(st.BasePath+st.RemoteFileName, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}