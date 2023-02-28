package engine

import (
	"log"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parserResult, err := worker(request)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v\n", item)
		}
	}
}
