package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anaskhan96/soup"
)

func ScrapeDHSKerala(c chan string) {
	resp, err := soup.Get("http://dhs.kerala.gov.in/%e0%b4%a1%e0%b5%86%e0%b4%af%e0%b4%bf%e0%b4%b2%e0%b4%bf-%e0%b4%ac%e0%b5%81%e0%b4%b3%e0%b5%8d%e0%b4%b3%e0%b4%b1%e0%b5%8d%e0%b4%b1%e0%b4%bf%e0%b4%a8%e0%b5%8d%e2%80%8d/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("h3", "class", "entry-title").Find("a")
	c<-"data/"+strings.ReplaceAll(links.Text(),"/","-")+".pdf"
}

func main() {
	files, err := filepath.Glob("data/*.pdf")
    if err != nil {
		fmt.Println(err)
		os.Exit(1)
    }

	c := make(chan string)
    for _, f := range files {
			go ScrapeDHSKerala(c)
			x := <-c
			if f == x {
				fmt.Println("The pdf file is already latest")
			} else {
				fmt.Printf("You need latest pdf file : %s(local) != %s(remote)\n",f,x)
			}
    }
	// fmt.Println(ScrapeDHSKerala())
}
