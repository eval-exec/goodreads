package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

var (
	c    *colly.Collector
	tags []string
)

const (
	baseUrl = "https://goodreads.com"
)

func allTags() (tags []string) {
	log.Println("on all tags")
	c.OnHTML("ul.listTagsTwoColumn", func(element *colly.HTMLElement) {
		element.ForEach("li.greyText", func(_ int, element2 *colly.HTMLElement) {
			element2.ForEach("a.gr-hyperlink", func(_ int, element3 *colly.HTMLElement) {
				href := element3.Attr("href")
				li := strings.LastIndex(href, "/")
				thistag := strings.TrimSpace(href[li+1:])
				if thistag != specificTag {
					return
				}

				tagurl := baseUrl + href
				log.Println(tagurl)
				if tagurl == "https://goodreads.com/quotes/tag/hope" {
					for page := 1; page <= 2; page++ {
						visitUrl := fmt.Sprintf("%s?page=%d", tagurl, page)
						c.Visit(visitUrl)
					}
				}
			})
		})
	})

	c.OnHTML("div.quoteDetails", func(element *colly.HTMLElement) {
		element.ForEach("div.quoteText", func(i int, element *colly.HTMLElement) {
			var content = element.Text
			const cdata = "//<![CDATA"
			if strings.Contains(content, cdata) {
				index := strings.Index(content, cdata)
				content = content[:index]
			}
			after := strings.SplitAfter(content, "  ―\n  ")
			quote := strings.TrimSpace(strings.TrimSuffix(after[0], "  ―\n  "))
			author := strings.TrimSpace(after[1])
			aus := strings.Split(author, ",")

			var footer string
			for _, au := range aus {
				footer += strings.TrimSpace(au)
			}

			csvW.Write([]string{
				quote, footer,
			})
			log.Println(quote)
			log.Println(footer)
		})

	})

	log.Println(c.Visit(baseUrl + "/quotes"))
	return
}
