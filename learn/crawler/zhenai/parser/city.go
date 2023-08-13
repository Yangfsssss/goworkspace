package parser

import (
	"engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParserResult {
	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])

		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
