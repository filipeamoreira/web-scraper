package main

import (
	"fmt"
	// "io"
	"log"
	"net/http"
	// "os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Running main")
	scrapeJobPage()
}

func scrapeJobPage() {
	response, err := http.Get("http://filipeamoreira.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find(".nav li").Each(func(index int, item *goquery.Selection) {
		// fmt.Printf(item.Text())
		link := item.Text()
		fmt.Printf("Link %d: %s\n", index, link)
	})

	// io.Copy(os.Stdout, response.Body)
}
