package v1

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

const (
	cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	URL         = "http://www.zhenai.com/zhenghun"
)

func ZhengAiWang() {
	resp, err := http.Get(URL)
	if err != nil {
		panic(fmt.Errorf("Error: http Get, err is %v\n", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: statuscode is ", resp.StatusCode)
		return
	}
	utf8Reader := transform.NewReader(resp.Body, determineEncoding(resp.Body).NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		fmt.Println("Error read body, error is ", err)
	}
	getAllCityUrl(body)
	fmt.Println("body is ", string(body))
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

// get all city and url
func getAllCityUrl(body [] byte) {
	compile := regexp.MustCompile(cityListReg)
	subMatch := compile.FindAllSubmatch(body, -1)
	for _, matches := range subMatch {
		//打印
		fmt.Printf("City:%s URL:%s\n", matches[2], matches[1])
	}
	fmt.Printf("Matches count: %d\n", len(subMatch))
}


