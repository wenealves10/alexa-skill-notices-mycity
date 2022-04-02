package main

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type News struct {
	Title        string `json:"title,omitempty"`
	Content      string `json:"content,omitempty"`
	NumberNotice int    `json:"number_notice,omitempty"`
}

func news(url string) []News {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)

	var news News
	var notices []News

	// On every a element which has href attribute call callback
	c.OnHTML(".featured-posts", func(e *colly.HTMLElement) {

		e.ForEach("article", func(i int, h *colly.HTMLElement) {
			news.Title = h.ChildText("div.entry-header > h2")
			news.Content = h.ChildText("div.entry-content")
			news.NumberNotice = i + 1
			notices = append(notices, news)
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://" + url)

	return notices
}

func getAllNews(c echo.Context) error {
	return c.JSON(http.StatusOK, news("www.blogdoacelio.com.br"))
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", getAllNews)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
