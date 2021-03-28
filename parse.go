package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type Course struct {
	Title       string
	Description string
	Author      string
	Url         string
}

func GetCourses(query string, limit int) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML(`a[data-track-component=search_card_title_link]`, func(e *colly.HTMLElement) {
		err := e.Request.Visit(e.Attr("href"))
		if err != nil {
			log.Fatal(err)
		}
	})

	c.OnHTML(`div[class=AboutCourse] div[class='rc-TogglableContent about-section collapsed'] p`, func(e *colly.HTMLElement) {
		c := Course{}
		c.Description = e.Text
	})

	// c.OnHTML(`a[class='label-text box arrow']`, func(e *colly.HTMLElement) {
	// 	if len(coursesCollected) >= limit {
	// 		return
	// 	}
	// 	log.Println("limit not reached")
	// 	err := e.Request.Visit(e.Attr("href"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// })

	c.Visit(fmt.Sprintf("https://www.coursera.org/courses?query=%s", query))

	c.OnScraped(func(r *colly.Response) {
		log.Println("Finished", r.Request.URL)
	})
}
