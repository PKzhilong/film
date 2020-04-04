package parse

import (
	"bytes"
	"filmspider/engine"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
)

func CategoryDetail(content []byte, db *gorm.DB, categoryID int, categories *engine.Categories) engine.ParseResult {

	bodyReader := bytes.NewReader(content)
	cont, err := goquery.NewDocumentFromReader(bodyReader)
	var parse engine.ParseResult
	if err != nil {
		return parse
	}

	cont.Find(".xing_vb ul").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		li := selection.Find("li")
		name := li.Find(".xing_vb4").Text()
		url, _ := li.Find(".xing_vb4 a").Attr("href")
		if len(name) < 1 || len(url) < 1 {
			return
		}

		request := engine.Request{
			Url:        host + url,
			ParserFunc: engine.NilParaser,
		}

		parse.Requests = append(parse.Requests, request)
		fmt.Printf("电影名： %s 地址: %s\n", name, url)
	})

	// 获取分页数据
	cont.Find(".pages a").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text != "下一页" {
			return
		}
		url,_ := selection.Attr("href")
		request := engine.Request{
			Url:        host + url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return CategoryDetail(bytes, db, categoryID, categories)
			},
		}
		parse.Requests = append(parse.Requests, request)
	})
	return parse

}