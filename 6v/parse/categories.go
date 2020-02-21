package parse

import (
	"filmspider/engine"
	"log"
	"regexp"
)

var categoryParse = regexp.MustCompile(`<li class="menu-item "><a href="([^"]+)">([^<]+)</a></li>`)

func CategoryParse(content []byte) engine.ParseResult {
	result := categoryParse.FindAllSubmatch(content, -1);
	var parse engine.ParseResult
	var requests []engine.Request
	var items []interface{}
	for _, v := range result {
		request := engine.Request{
			Url: string(v[1]),
			ParserFunc: CategoryDetail,
		}
		item := engine.Item{
			Name: string(v[2]),
			Url: string(v[1]),
			Type: 1,
		}

		log.Printf("分类名：%s Url：%s\n", item.Name, item.Url)
		requests = append(requests, request)
		items = append(items, item)
	}
	parse.Requests = requests
	parse.Items = items

	return parse
}
