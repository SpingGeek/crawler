package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href != "index.html" {
			c.Visit(e.Request.AbsoluteURL(href))
		}
	})

	c.OnHTML(".article-title", func(h *colly.HTMLElement) {
		title := h.Text
		fmt.Printf("title: %v\n", title)
	})

	c.OnHTML(".article", func(h *colly.HTMLElement) {
		content, _ := h.DOM.Html()
		fmt.Printf("content: %v\n", content)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://gorm.io/zh_CN/docs/")
}
