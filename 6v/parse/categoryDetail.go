package parse

import (
	"filmspider/engine"
	"log"
	"regexp"
)


var detail = regexp.MustCompile(`<a href="([^"]+)" class="zoom" rel="bookmark" title="([^"]+)">[^<]*<img src="([^"]+)" alt="[^"]+" />[^<]*</a>`)

var nextPage = regexp.MustCompile(`<a href="([^"]+)" class="next">下一页</a>`)

func CategoryDetail(content []byte) engine.ParseResult  {
	result := detail.FindAllSubmatch(content, -1)

	var parse engine.ParseResult
	var requests []engine.Request
	var items []interface{}

	nextPageRes := nextPage.FindAllSubmatch(content, 1)
	if len(nextPageRes) > 0 {
		request := engine.Request{
			Url: string(nextPageRes[0][1]),
			ParserFunc: CategoryDetail,
		}
		requests = append(requests, request)
	}

	for _, v := range result {
		fileName := string(v[2])
		request := engine.Request{
			Url:        string(v[1]),
			//ParserFunc: FilmDetail,
			//ParserFunc: func(c []byte) engine.ParseResult {
			//	return FilmDetail(c, string(v[1]))
			//},
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return FilmDetailByDuc(bytes, fileName)
			},
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
