package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type News struct {
	Title        string `json:"title,omitempty"`
	Content      string `json:"content,omitempty"`
	NumberNotice int    `json:"number_notice,omitempty"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.blogdoacelio.com.br"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".featured-posts", func(e *colly.HTMLElement) {

		var news News
		var notices []News

		e.ForEach("article", func(i int, h *colly.HTMLElement) {
			news.Title = h.ChildText("div.entry-header > h2")
			news.Content = h.ChildText("div.entry-content")
			news.NumberNotice = i + 1
			notices = append(notices, news)
		})
		jsonClient, err := json.Marshal(notices)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		fmt.Println(string(jsonClient))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.blogdoacelio.com.br")
}
