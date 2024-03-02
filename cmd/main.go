package main

import (
	"fmt"
	"log"
	"web-scrapper-go/cmd/cli"
	"web-scrapper-go/scrapper"
)

func main() {
	fmt.Println("start Maktabkhooneh Downloader")

	args := cli.Execute()

	c, err := scrapper.Config(args)

	if err != nil {
		log.Fatalf("bad args %s", err)
	}
	_, err = scrapper.FetchCourse(c)
	if err != nil {
		log.Fatalf("fetch course failure: %s", err)
	}

	// visiting the first page
	err = c.Visit(args.Url)
	c.Wait()

	if err != nil {
		log.Fatalf("site visiting err:%s", err)
	}

}
