package parse

import (
	"filmspider/engine"
	"log"
	"regexp"
)

var coverDetail = regexp.MustCompile(`<p>[^<]*<img alt[=|"]* src="([^"]+)"[^>]*>`)

func FilmDetail(content []byte) engine.ParseResult{
	result := coverDetail.FindAllSubmatch(content, -1)

	var parse engine.ParseResult
	var requests []engine.Request
	var items []interface{}

	coverImg := ""
	infoImg := ""
	for k, v := range result {
		if k == 1 {
			infoImg = string(v[1])
			continue
		}

		coverImg = string(v[1])
		request := engine.Request{
			Url: string(v[1]),
			ParserFunc: engine.NilParaser,
		}
		requests = append(requests, request)
	}
	log.Printf("封面图：%s 详情图：%s", coverImg, infoImg)
	parse.Requests = requests
	parse.Items = items
	return engine.ParseResult{}
}