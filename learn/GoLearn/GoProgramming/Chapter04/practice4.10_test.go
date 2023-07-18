package Chapter03

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

func start0410() {
	result, err := SearchIssues(os.Args[4:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s  %s\n",
			item.Number, item.User.Login, item.Title, item.State)
	}
}

func start04102() {
	result, err := SearchIssues(os.Args[4:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	//当前时间
	now := time.Now().Unix()
	//前一个月
	preMonth := now - 30*24*3600
	//前一年
	preYear := now - 365*24*3600

	var notMonth []*Issue
	var notYear []*Issue
	var overYear []*Issue

	for _, item := range result.Items {
		createTime := item.CreatedAt.Unix()
		if createTime > preMonth {
			notMonth = append(notMonth, item)
			continue
		}
		if createTime < preMonth && createTime > preYear {
			notYear = append(notYear, item)
			continue
		}
		overYear = append(overYear, item)
	}

	fmt.Println("issues(不到一个月):")
	for _, item := range notMonth {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Println("issues(不到一年):")
	for _, item := range notYear {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Println("issues(超过一年):")
	for _, item := range overYear {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

// go test -v -run Test0410 practice4.10_test.go  -args go java redis
func Test0410(t *testing.T) {
	//start0410()
	start04102()
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	url := IssuesURL + "?q=" + q
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

type IssuesSearchResult struct {
	TotalCount int `jsontest:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `jsontest:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `jsontest:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `jsontest:"html_url"`
}
