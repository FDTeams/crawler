package engine

import (
	"log"

	"github.com/seaung/crawler/pkg/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		log.Printf("Fetcher %s\n", request.Url)
		content, err := fetcher.Fetcher(request.Url)
		if err != nil {
			continue
		}

		parserResult := request.ParseFunc(content)

		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item %v\n", item)
		}
	}
}
