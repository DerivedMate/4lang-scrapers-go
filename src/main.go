package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

func getLinks(url string, ch chan []string) {
	var urls []string

	c := colly.NewCollector()

	// -------- Events -------- //
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("img", func(el *colly.HTMLElement) {
		urls = append(urls, el.Attr("src"))
	})

	c.OnScraped(func(_ *colly.Response) {
		ch <- urls
	})

	// -------- Initialization -------- //
	err := c.Visit(url)

	if err != nil {
		panic(err)
	}
}

func makeFile(name string) *os.File {
	f, err := os.Create("./images/" + name)
	if err != nil {
		log.Fatalf("Error creating file %s", err)
	}
	return f
}

func makeFileName(n int, ext string) string {
	s := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		s += string(rand.Uint32())
	}

	return s + ext
}

func makeHTTPReq(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting the url \"%s\": %v", url, err)
	}

	return resp
}

func downloadImages(src string, links []string) {
	extRe := regexp.MustCompile(`(\.\w{3})`)
	nameRe := regexp.MustCompile(`([\w\d-_]{5,})`)
	for _, link := range links {
		extMatches := extRe.FindAllString(link, 3)
		ext := extMatches[2]

		fnameMatches := nameRe.FindAllString(link, 5)
		fname := fnameMatches[3]
		if fname == "" || ext == "" {
			fmt.Printf("Skipping %s\n", link)
			continue
		}

		f := makeFile(fname + ext)
		resp := makeHTTPReq(link)

		defer resp.Body.Close()
		defer f.Close()

		size, err := io.Copy(f, resp.Body)
		if err != nil {
			log.Fatalf("Error writing to \"%s\": %v", fname, err)
		}

		fmt.Printf("Downloaded a file \"%s\" of size: %v\n", fname, size)
	}

	fmt.Printf("----- Finized %s ------\n\n\n", src)
}

func main() {
	toScrape := []string{"https://www.deviantart.com/?offset=0", "https://www.deviantart.com/?offset=240"}
	ch := make(chan []string)

	for _, l := range toScrape {
		go getLinks(l, ch)
		links := <-ch

		fmt.Printf("Found %v links : %s\n", len(links), l)
		downloadImages(l, links)
	}

}
