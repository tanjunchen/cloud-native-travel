package v2

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    [] interface{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	PayLoad interface{}
}

type Profile struct {
	Name     string
	Gender   string
	Age      int
	Height   int
	Weight   int
	Income   string
	Marriage string
	Address  string
}

func Start() {
	request := Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: ParseCityList,
	}
	Run(request)
}

///////////////engine////////////
func Run(seeds ...Request) {
	var requestsQueue []Request
	requestsQueue = append(requestsQueue, seeds...)
	for len(requestsQueue) > 0 {
		r := requestsQueue[0]
		requestsQueue = requestsQueue[1:]
		log.Printf("fetching url:%s\n", r.Url)

		body, err := Fetch(r.Url)
		if err != nil {
			log.Printf("fetch url: %s; err: %v\n", r.Url, err)
			continue
		}
		result := r.ParserFunc(body)
		requestsQueue = append(requestsQueue, result.Requests...)
		for _, item := range result.Items {
			log.Printf("get item is %v\n", item)
		}
	}
}

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest is err ", err)
		return nil, fmt.Errorf("NewRequest is err %v\n", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error: http Get, err is %v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: StatusCode is %d\n", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(bodyReader, determineEncoding(bodyReader).NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

////////////////Parser///////////////////
var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)kg</div>`)

var incomeRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
var addressRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>工作地:([^<]+)</div>`)

func ParserProfile(content []byte, name string, gender string) ParserResult {
	profile := Profile{}
	profile.Name = name
	profile.Gender = gender
	if age, err := strconv.Atoi(extractString(content, ageRe)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(content, heightRe)); err == nil {
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractString(content, weightRe)); err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(content, incomeRe)
	profile.Marriage = extractString(content, marriageRe)
	profile.Address = extractString(content, addressRe)

	result := ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	subMatch := re.FindSubmatch(contents)
	if len(subMatch) >= 2 {
		return string(subMatch[1])
	} else {
		return ""
	}
}

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 解析城市列表
func ParseCityList(bytes []byte) ParserResult {
	re := regexp.MustCompile(cityListRe)
	subMatch := re.FindAllSubmatch(bytes, -1)
	result := ParserResult{}
	limit := 10
	for _, item := range subMatch {
		result.Items = append(result.Items, "City:"+string(item[2]))
		result.Requests = append(result.Requests, Request{
			Url:        string(item[1]),
			ParserFunc: ParserCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

// 城市页面用户解析器
func ParserCity(bytes []byte) ParserResult {
	subMatch := cityRe.FindAllSubmatch(bytes, -1)
	genderMatch := sexRe.FindAllSubmatch(bytes, -1)

	result := ParserResult{}

	for k, item := range subMatch {
		name := string(item[2])
		gender := string(genderMatch[k][1])

		result.Items = append(result.Items, "User:"+name)
		result.Requests = append(result.Requests, Request{
			Url: string(item[1]),
			ParserFunc: func(bytes []byte) ParserResult {
				return ParserProfile(bytes, name, gender)
			},
		})
	}
	return result
}

func determineEncoding(r io.Reader) encoding.Encoding {
	//这里的r读取完得保证resp.Body还可读
	body, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		fmt.Println("Error: peek 1024 byte of body err is ", err)
	}

	//这里简化,不取是否确认
	e, _, _ := charset.DetermineEncoding(body, "")
	return e
}