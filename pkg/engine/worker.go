package engine

import (
	"log"

	"github.com/seaung/crawler/pkg/fetcher"
)

func worker(request Request) (ParseResult, error) {
	log.Printf("Fetcher %s\n", request.Url)
	content, err := fetcher.Fetcher(request.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return request.ParseFunc(content), nil
}
