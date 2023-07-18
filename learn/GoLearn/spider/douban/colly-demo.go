package main

import (
	"encoding/csv"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func demo01() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Visited", response.Request.URL)
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
}

func demo02() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "*.douban.*",
		Delay:        0,
		RandomDelay:  0,
		Parallelism:  5,
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
}

func demo03() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 5})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal(err)
		}
		nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)
		for _, node := range nodes {
			url := htmlquery.FindOne(node, "./a/@href")
			title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
			log.Println(strings.Split(htmlquery.InnerText(url), "/")[4],
				htmlquery.InnerText(title))
		}
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
}

func fetchUrl(url string) *html.Node {
	log.Println("fetch url : ", url)
	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseUrl(url string, ch chan bool, f *os.File) {
	doc := fetchUrl(url)
	nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)
	for _, node := range nodes {
		url := htmlquery.FindOne(node, "./a/@href")
		title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
		_, err := f.WriteString(strings.Split(htmlquery.InnerText(url), "/")[4] + "\t" +
			htmlquery.InnerText(title) + "\n")
		checkError(err)
	}
	ch <- true
}

func test01() {
	start := time.Now()
	ch := make(chan bool)
	f, err := os.Create("spider/douban/movie.txt")
	checkError(err)
	defer f.Close()
	_, err = f.WriteString("ID\tTitle\n")
	checkError(err)
	for i := 0; i < 10; i++ {
		go parseUrl("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch, f)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	f.Sync()
	log.Printf("Took %s", time.Since(start))
}

func parseCSVUrl(url string,ch chan bool,w * csv.Writer)  {
	doc := fetchUrl(url)
	nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)
	for _, node := range nodes {
		url := htmlquery.FindOne(node, "./a/@href")
		title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
		err := w.Write([]string{
			strings.Split(htmlquery.InnerText(url), "/")[4],
			htmlquery.InnerText(title)})
		checkError(err)
	}
	ch <- true
}

func test02()  {
	start := time.Now()
	ch := make(chan bool)
	f, err := os.Create("spider/douban/movie.csv")
	checkError(err)
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	err = writer.Write([]string{"ID", "Title"})
	checkError(err)
	for i := 0; i < 10; i++ {
		go parseCSVUrl("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch, writer)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	log.Printf("Took %s", time.Since(start))
}

func main() {
	test02()
}
