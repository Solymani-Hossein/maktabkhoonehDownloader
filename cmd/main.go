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
	Seasons []Season
}

type Season struct {
	Title string
	Links []map[string]string
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
	//
	// main of page element
	mainElement := ".w-full.md\\:mt-20"
	// name of course element
	seasonHeadElement := ".course-page__title"
	// title of evry season element
	titleElement := ".base-chapter__title"
	//main of season element
	seasonMainElement := ".first\\:rounded-t:nth-child(%d) .p-10"

	c.OnHTML(mainElement, func(e *colly.HTMLElement) {

		var season Season

		//get course name
		courseName := e.ChildText(seasonHeadElement)
		course.Name = courseName

		// loop in title of season
		e.ForEachWithBreak(titleElement, func(index int, h *colly.HTMLElement) bool {
			currentIndex := index + 1
			season.Title = h.Text

			seasonMain := fmt.Sprintf(seasonMainElement, currentIndex)
			e.ForEachWithBreak(seasonMain, func(x int, j *colly.HTMLElement) bool {
				title := j.ChildAttr("span.text-blue", "title")
				link := j.Attr("href")

				if title != "" {
					season.Links = append(season.Links, map[string]string{
						"title": title,
						"link":  link,
					})
				}
				return true
			})
			course.Seasons = append(course.Seasons, season)
			return true
		})

	})

	// visiting the first page
	err := c.Visit(urls)
	c.Wait()

	if err != nil {
		log.Fatalf("site visiting err:%s", err)
	}

	// opening the CSV file
	file, err := os.Create(fmt.Sprintf("%s.csv", course.Name))

	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}

	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{}

	for _, season := range course.Seasons {
		headers = append(headers, season.Title)
		for _, subTitle := range season.Links {
			fmt.Println(subTitle["title"])
			fmt.Println(subTitle["link"])
			fmt.Println("==================================================")
		}
	}

	writer.Write(headers)

	//for _, data := range course.Seasons {
	//	for i, sub := range data.Links {
	//		subtitles := []string{
	//			sub[i],
	//		}
	//		writer.Write(subtitles)
	//	}
	//}

	defer writer.Flush()

}
