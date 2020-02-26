package parse

import (
	"filmspider/engine"
	"filmspider/model"
	"filmspider/repository"
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
		item := model.Category {
			Name: string(v[2]),
		}

		//判断分类是否已经则不进行添加入库
		newC := repository.Category{}.GetCategoryByName(item.Name)
		ID := newC.ID
		if newC.ID == 0 {
			repository.Category{}.CreateCategory(&item)
			ID = item.ID
		}
		request := engine.Request{
			Url: string(v[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return CategoryDetail(bytes, ID)
			},
		}

		log.Printf("分类名：%s Url：%s\n", item.Name, v[1])
		requests = append(requests, request)
		items = append(items, item)
	}
	parse.Requests = requests
	parse.Items = items

	return parse
}
