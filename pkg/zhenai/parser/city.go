package parser

import (
	"regexp"

	"github.com/seaung/crawler/pkg/engine"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

func ParseCity(content []byte) engine.ParseResult {
	submatch := cityRe.FindAllSubmatch(content, -1)
	gendermatch := sexRe.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for k, item := range submatch {
		name := string(item[2])
		gender := string(gendermatch[k][1])
		result.Items = append(result.Items, "User : "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(item[1]),
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name, gender)
			},
		})
	}
	return result
}
