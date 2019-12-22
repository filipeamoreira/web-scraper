package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Running main")
	scrapeUpworkJob()
}

type Project struct {
	title           string
	category        string
	description     string
	time_commitment string
	length          string
	expert_level    string
}

func scrapeUpworkJob() {
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

	project.time_commitment = jobFeatures[0]
	project.length = jobFeatures[1]
	project.expert_level = jobFeatures[2]

	persistJob(project)
	// fmt.Println(project)
}

func persistJob(project Project) {
	setupDB(project)
}

func setupDB(project Project) {
	connStr := "postgresql://scrape:scrape@localhost/scrape?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB setup correctly")
	// fmt.Println(project)
	insertSqlStatement := `
        INSERT INTO jobs (title, category, description, time_commitment, length, expert_level)
        VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = db.Exec(insertSqlStatement, project.title, project.category, project.description, project.time_commitment, project.length, project.expert_level)

	if err != nil {
		log.Fatal(err)
	}

	select_rows, err := db.Query(`SELECT * FROM jobs`)
	fmt.Println(select_rows)
}
