package main

import (
	"github.com/antchfx/htmlquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Movie struct {
	ID    int    `gorm:"AUTO_INCREMENT"`
	Title string `gorm:"type:varchar(100);unique_index"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchUrl(url string) *html.Node {
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

func parseUrl(url string, ch chan bool, db *gorm.DB) {
	doc := fetchUrl(url)
	nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)

	for _, node := range nodes {
		url := htmlquery.FindOne(node, "./a/@href")
		title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)

		id, _ := strconv.Atoi(strings.Split(htmlquery.InnerText(url), "/")[4])

		movie := &Movie{
			ID:    id,
			Title: htmlquery.InnerText(title),
		}
		db.Create(&movie)
	}

	ch <- true
}

func demo01() {
	start := time.Now()
	ch := make(chan bool)
	db, err := gorm.Open("mysql", "root:a@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	checkError(err)
	db.DropTableIfExists(&Movie{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Movie{})
	for i := 0; i < 11; i++ {
		go parseUrl("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch, db)
	}
	for i := 0; i < 11; i++ {
		<-ch
	}
	log.Printf("Took %s", time.Since(start))
}

func findMovies() {
	db, err := gorm.Open("mysql", "root:a@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	checkError(err)

	var movie Movie
	var movies []Movie

	db.First(&movie, 30170448)
	log.Println(movie)
	log.Println(movie.ID, movie.Title)

	db.Order("id").Limit(3).Find(&movies)
	log.Println(movies)
	log.Println(movies[0].ID)

	db.Order("id desc").Limit(3).Offset(1).Find(&movies)
	log.Println(movies)

	db.Select("title").Find(&movies, 30170448)
	log.Println(movies)

	db.Select("title").First(&movies, "title = ?", "四个春天")
	log.Println(movie)

	var count int64
	db.Where("id = ?", 30170448).Or("title = ?", "四个春天").Find(&movies).Count(&count)
	log.Println(count)
}

func main() {
	//demo01()
	findMovies()
}
