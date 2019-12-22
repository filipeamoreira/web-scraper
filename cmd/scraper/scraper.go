package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Running main")
	scrapeJobPage()
}

type Project struct {
	title          string
	category       string
	description    string
	timeCommitment string
	length         string
	expertLevel    string
}

func scrapeJobPage() {
	response, err := http.Get("https://www.upwork.com/job/Back-end-developer-Ruby-Rails_~015192b45cd3e4eb34/")
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

	project := Project{}

	project.title = document.Find("h2.m-0-bottom").First().Text()
	project.category = document.Find("a.specialization").First().Text()
	project.description = document.Find("div.job-description").First().Text()
	jobFeatures := []string{}

	document.Find("ul.job-features li").Each(func(index int, item *goquery.Selection) {
		jobFeatures = append(jobFeatures, item.Text())
	})

	project.timeCommitment = jobFeatures[0]
	project.length = jobFeatures[1]
	project.expertLevel = jobFeatures[2]

	fmt.Println(project)
}
