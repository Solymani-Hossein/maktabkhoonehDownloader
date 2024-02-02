package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type Course struct {
	Name    string
	Seasons []map[int]Season
}

type Season struct {
	Titles []map[int]string
	Links  []map[string]string
}

func main() {

	var course Course
	// initializing the slice of structs that will contain the scraped data

	// the first pagination URL to scrape
	urls := "https://maktabkhooneh.org/course/آموزش-تحلیل-بدافزار-مقدماتی-mk933/"

	// initializing a Colly instance

	c := colly.NewCollector(
		colly.MaxDepth(30),
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		RandomDelay: 2 * time.Second,
	})

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	// iterating over the list of pagination links to implement the crawling logic
	mainElement := ".w-full.md\\:mt-20"
	seasonHeadElement := ".course-page__title"
	titleElement := ".base-chapter__title"
	seasonMainElement := ".first\\:rounded-t:nth-child(%d) .p-10"

	c.OnHTML(mainElement, func(e *colly.HTMLElement) {

		var season Season

		courseName := e.ChildText(seasonHeadElement)
		course.Name = courseName

		e.ForEachWithBreak(titleElement, func(i int, h *colly.HTMLElement) bool {

			season.Titles = append(season.Titles, map[int]string{
				i + 1: h.Text,
			})

			seasonMain := fmt.Sprintf(seasonMainElement, i+1)
			e.ForEachWithBreak(seasonMain, func(x int, j *colly.HTMLElement) bool {

				season.Links = append(season.Links, map[string]string{
					j.Text: j.Attr("href"),
				})
				return true
			})
			course.Seasons = append(course.Seasons, map[int]Season{
				i + 1: season,
			})
			return true
		})

	})

	fmt.Println("course Name:", course.Name)
	for _, data := range course.Seasons {
		fmt.Println(data)
	}
	// visiting the first page
	err := c.Visit(urls)
	c.Wait()

	if err != nil {
		log.Fatalf("site visiting err:%s", err)
	}

	// opening the CSV file
	file, err := os.Create("products.csv")

	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}

	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	//headers := []string{}

	//for _, data := range  {
	//	headers = append(headers, data.title)
	//}

	//writer.Write(headers)

	// writing each Pokemon product as a CSV row
	//for _, data := range dataStructures {
	//	for _, sub := range data.subtitle {
	//		subtitles := []string{
	//			sub,
	//		}
	//		writer.Write(subtitles)
	//	}
	//}

	defer writer.Flush()

}
