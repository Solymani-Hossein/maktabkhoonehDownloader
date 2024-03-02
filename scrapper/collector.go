package scrapper

import (
	"fmt"
	"log"
	"web-scrapper-go/internal/entity"
	"web-scrapper-go/internal/utils"

	"github.com/gocolly/colly"
)

func FetchCourse(c *colly.Collector) (course entity.Course, err error) {

	// main of page element
	mainElement := ".w-full.md\\:mt-20"
	// name of course element
	seasonHeadElement := ".course-page__title"
	// title of evry season element
	titleElement := ".base-chapter__title"
	//main of season element
	seasonMainElement := ".first\\:rounded-t:nth-child(%d) .p-10"

	c.OnHTML(mainElement, func(main *colly.HTMLElement) {

		//get course name
		course.Name = main.ChildText(seasonHeadElement)

		// loop in title of season
		main.ForEachWithBreak(titleElement, func(index int, seasonElement *colly.HTMLElement) bool {
			var season entity.Season
			currentIndex := index + 1
			season.Title = seasonElement.Text

			seasonMain := fmt.Sprintf(seasonMainElement, currentIndex)
			main.ForEachWithBreak(seasonMain, func(x int, sm *colly.HTMLElement) bool {
				isHref := sm.Attr("href")

				if isHref != "" {
					domain := "https://maktabkhooneh.org"
					href := fmt.Sprintf("%s%s", domain, isHref)
					fmt.Println(href)

					link, err := GetLink(c, href)
					if err != nil {
						fmt.Println("failure to get url of video")
					}

					link.Name = sm.ChildAttr("span.text-blue", "title")
					season.Links = append(season.Links, link)
				}

				return true
			})
			course.Seasons = append(course.Seasons, season)
			return true
		})
		PreparLinks(course)
	})
	return course, nil
}

func GetLink(c *colly.Collector, href string) (link entity.Link, err error) {

	DownloadLinkButtonElement := ".js-copy-Popup+ .unit-content--download .button--round"

	c.OnHTML(DownloadLinkButtonElement, func(dl *colly.HTMLElement) {
		link.Href = dl.Attr("href")
	})

	// visiting the first page
	err = c.Visit(href)

	if err != nil {
		log.Fatalf("site visiting for get video link err:%s", err)
	}
	return

}

func PreparLinks(course entity.Course) {

	for _, dir := range course.Seasons {
		path := fmt.Sprintf("%s/%s", course.Name, dir.Title)
		err := utils.CreateDirectory(path)
		if err != nil {
			fmt.Printf("create dir failuer. %s", err)
		}

		for _, link := range dir.Links {
			linkPath := fmt.Sprintf("%s/%s", path, link.Name)
			err = utils.DownloadFile(linkPath, link.Href)
			if err != nil {
				fmt.Printf("downlaod file failuer. %s", err)
			}
		}
	}
}
func Download() {}
