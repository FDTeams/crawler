package parser

import (
	"regexp"

	"github.com/seaung/crawler/pkg/engine"
)

const cityListRE = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRE)

	submatch := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, item := range submatch {
		result.Items = append(result.Items, "City: "+string(item[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
