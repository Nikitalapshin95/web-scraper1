package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Industry struct {
	Url, Image, Name string
}

func main() {
	file_name := "auto.csv"
	file, err := os.Create(file_name)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Price", "Link"})

	c := colly.NewCollector(colly.AllowedDomains("auto.ru"))

	c.OnHTML(".ListingItem", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		price := e.ChildText(".ListingItemPrice__content")
		title := e.ChildText(".ListingItemTitle__link")

		writer.Write([]string{
			title,
			price,
			link,
		})
	})

	c.Visit("https://auto.ru/moskva/cars/all/")

	visited := 0
	c.OnRequest(func(r *colly.Request) {
		if visited > 100 {
			r.Abort()
		}
		visited++
	})
}
