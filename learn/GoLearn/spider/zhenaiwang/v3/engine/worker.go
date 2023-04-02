package engine

import (
	"go-learn/spider/zhenaiwang/v3/fetcher"
	"log"
)

func worker(request Request)(ParseResult, error)  {
	log.Printf("Fetching %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch error, Url: %s %v\n", request.Url, err)
		return ParseResult{}, err
	}
	return request.ParseFunc(content), nil
}
