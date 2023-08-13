package parser

import (
	"engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParserResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		result.Items = append(result.Items, "City "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseCity,
		})
	}

	return result
}
