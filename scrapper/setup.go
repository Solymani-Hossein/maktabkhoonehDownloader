package scrapper

import (
	"log"
	"net/http"
	"time"
	"web-scrapper-go/internal/entity"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func Config(args entity.Args) (*colly.Collector, error) {

	// initializing a Colly instance
	c := colly.NewCollector(
		colly.MaxDepth(30),
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		cookie := "sessionid=m5cm2tqg50ohxlayly628myjrpa3zcfl"
		r.Headers.Set("cookie", cookie)
	})

	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 4,
		RandomDelay: 2 * time.Second,
	})

	if err != nil {
		log.Printf("error in LimitRule. %s", err)
	}

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	return c, nil

}
