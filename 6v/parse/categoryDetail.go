package parse

import (
	"filmspider/engine"
	"log"
	"regexp"
)


var detail = regexp.MustCompile(`<a href="([^"]+)" class="zoom" rel="bookmark" title="([^"]+)">[^<]*<img src="([^"]+)" alt="[^"]+" />[^<]*</a>`)

func CategoryDetail(content []byte) engine.ParseResult  {
	result := detail.FindAllSubmatch(content, -1)

	var parse engine.ParseResult
	var requests []engine.Request
	var items []interface{}

	for _, v := range result {
		request := engine.Request{
			Url:        string(v[1]),
			ParserFunc: FilmDetail,
		}
		item := engine.Item{
			Url: string(v[1]),
			Name: string(v[2]),
			Type: 2,
		}
		requests = append(requests, request)
		items = append(items, item)
		log.Printf("片名：%s 详情地址： %s", item.Name, item.Url)
	}

	parse.Requests = requests
	parse.Items = items
	return parse
}
