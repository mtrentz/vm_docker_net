package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getURLs(src string) {
	r, err := http.Get(src)
	if err != nil {
		return
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		fmt.Println(r.StatusCode)
	}

	d, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// var urls []string
	d.Find("a").Each(func(_ int, a *goquery.Selection) {
		h, _ := a.Attr("href")
		fmt.Println(h)
	})
}

func main() {
	getURLs("https://www.maxiquim.com.br/")
}
