package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	var youtubeLinks []string

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		links := strings.Split(fmt.Sprintf("%s", e.Attr("href")), "\n")
		for _, url := range links {
			if strings.HasPrefix(url, "https://www.youtube.com") {
				if strings.Contains(url, "playlist") || strings.Contains(url, "blogs") {
					continue
				}
				youtubeLinks = append(youtubeLinks, url[:43]) // avoid ty links with ?s=131 postfix
			}
		}
	})

	for _, pageNumber := range []string{"0", "1", "2", "3", "4", "5"} {
		url := "https://www.tinymixtapes.com/features/2010s-favorite-100-music-releases-decade?page=" + pageNumber
		c.Visit(url)
	}

	f, err := os.Create("tinymixtapes-urls.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	for _, link := range youtubeLinks {
		fmt.Fprintln(f, link)
	}
}
