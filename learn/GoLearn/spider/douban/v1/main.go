package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/anaskhan96/soup"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func fetch(url string) string {
	fmt.Println("fetch url .... ", url)
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("http get error: ", err)
		return ""
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("http status code: ", response.StatusCode)
		return ""
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read error: ", err)
		return ""
	}
	return string(body)
}

func parse(url string) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := rp.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		fmt.Println(idRe.FindStringSubmatch(item[1])[1],
			titleRe.FindStringSubmatch(item[1])[1])
	}
}

func test01() {
	url := "https://movie.douban.com/top250?start=%s"
	start := time.Now()
	for i := 0; i < 10; i++ {
		parse(fmt.Sprintf(url, strconv.Itoa(25*i)))
	}
	fmt.Printf("Took Time %s", time.Since(start))
}

func parse2(url string, ch chan bool) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := rp.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		fmt.Println(idRe.FindStringSubmatch(item[1])[1],
			titleRe.FindStringSubmatch(item[1])[1])
	}
	ch <- true
}

func test02() {
	url := "https://movie.douban.com/top250?start=%s"
	start := time.Now()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go parse2(fmt.Sprintf(url, strconv.Itoa(25*i)), ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Printf("Took Time %s", time.Since(start))
}

func test03() {
	url := "https://movie.douban.com/top250?start=%s"
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			parse(fmt.Sprintf(url, strconv.Itoa(25*i)))
		}(i)
	}

	wg.Wait()
	fmt.Printf("Took %s", time.Since(start))
}

func fetchUrl(url string) *goquery.Document {
	fmt.Println("fetch url .... ", url)
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("http get error: ", err)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("http status code: ", response.StatusCode)
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return doc
}

func parseUrl(url string) {
	doc := fetchUrl(url)
	doc.Find("ol.grid_view li").Find(".hd").Each(func(index int, ele *goquery.Selection) {
		movieUrl, _ := ele.Find("a").Attr("href")
		fmt.Println(strings.Split(movieUrl, "/")[4], ele.Find(".title").Eq(0).Text(),
			ele.Find(".title").Eq(1).Text())
	})
}

func test04() {
	url := "https://movie.douban.com/top250?start=%s"
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			parseUrl(fmt.Sprintf(url, strconv.Itoa(25*i)))
		}(i)
	}

	wg.Wait()
	fmt.Printf("Took %s", time.Since(start))
}

func fetchSoup(url string) soup.Root {
	fmt.Println("fetch Url ", url)
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	}
	source, err := soup.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	doc := soup.HTMLParse(source)
	return doc
}

func parseSoup(url string, ch chan bool) {
	doc := fetchSoup(url)
	for _, root := range doc.Find("ol", "class", "grid_view").FindAll("div", "class", "hd") {
		movieUrl, _ := root.Find("a").Attrs()["href"]
		title := root.Find("span", "class", "title").Text()
		fmt.Println(strings.Split(movieUrl, "/")[4], title)
	}
	ch <- true
}

func test05() {
	url := "https://movie.douban.com/top250?start=%s"
	start := time.Now()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go parseSoup(fmt.Sprintf(url, strconv.Itoa(25*i)), ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Printf("Took %s", time.Since(start))
}

func fetchHtmlQuery(url string) *html.Node {
	log.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func parseHtmlQuery(url string, ch chan bool) {
	doc := fetchHtmlQuery(url)
	nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)
	for _, node := range nodes {
		url := htmlquery.FindOne(node, "./a/@href")
		title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
		log.Println(strings.Split(htmlquery.InnerText(url), "/")[4],
			htmlquery.InnerText(title))
	}
	ch <- true
}

func test06() {
	url := "https://movie.douban.com/top250?start=%s"
	start := time.Now()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go parseHtmlQuery(fmt.Sprintf(url, strconv.Itoa(25*i)), ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Printf("Took %s", time.Since(start))
}

func main() {
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	test06()
}
