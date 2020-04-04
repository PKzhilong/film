package parse

import (
	"bytes"
	"filmspider/engine"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"strings"
)

var host = "http://www.okzyw.com"

func Category(content []byte, d *gorm.DB, categories *engine.Categories) engine.ParseResult {

	bodyReader := bytes.NewReader(content)
	cont, err := goquery.NewDocumentFromReader(bodyReader)
	var parse engine.ParseResult
	if err != nil {
		return parse
	}

	cont.Find("#sddm li").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		selection.Find("a").Each(func(index int, se *goquery.Selection) {
			if index == 0 {
				name := se.Text()
				name = strings.Replace(name, "片", "", -1)
				src, _ := se.Attr("href")
				if name == "福利" || name == "解说" {
					return
				}

				url := host + src
				//获取内容分类
				contentID := 0
				for _, v := range categories.ContentTypes {
					if v.Name == name {
						contentID = v.ID
					}
				}

				request := engine.Request{
					Url: url,
					ParserFunc: func(bytes []byte) engine.ParseResult {
						return CategoryDetail(bytes, d, contentID, categories)
					},
				}
				parse.Requests = append(parse.Requests, request)
				fmt.Printf("%s %s\n", name, url)
				return
			}
		})

	})
	return parse
}
