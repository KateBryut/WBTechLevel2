package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

var cfg Config

type Config struct {
	url *url.URL
}

func (cfg *Config) parse() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "url missing")
		os.Exit(2)
	}
	var err error
	cfg.url, err = url.Parse(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't parce url: %s\n", err)
		os.Exit(2)
	}
}

func saveFile(p string, b []byte) error {
	p = strings.Replace(p, "/", "_", -1)
	if _, err := os.Stat("site/" + p); err == nil {
		return nil
	}
	if err := os.WriteFile("site/"+p, b, 0644); err != nil {
		return err
	}
	return nil
}

func wget(cfg *Config) {
	c := colly.NewCollector(
		colly.AllowedDomains(cfg.url.Host),
		colly.Async(true),
	)

	links := make(map[string]string)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if _, ok := links[link]; ok {
			return
		}
		links[link] = link
		// visit link found on page
		// only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		u := path.Join(r.Request.URL.Host, r.Request.URL.Path)

		if err := saveFile(u, r.Body); err != nil {
			log.Println(err)
		}
	})

	if err := c.Visit(cfg.url.String()); err != nil {
		log.Println(err)
	}
	c.Wait()
}

func main() {
	cfg.parse()
	wget(&cfg)
}
