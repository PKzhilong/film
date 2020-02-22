package parse

import (
	"bytes"
	"filmspider/engine"
	"filmspider/model"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
)

var film = map[int]string{
	1: "导演",
	2: "编剧",
	3: "主演",
	4: "类型",
	5: "地区",
	6: "语言",
	7: "字幕",
	8: "上映日期",
	9: "剧情介绍",
}

var coverDetail = regexp.MustCompile(`<p>[^<]*<img alt[=|"]* src="([^"]+)"[^>]*>`)
var filmInfo = regexp.MustCompile(`导演:(.+)?编剧:(.+)?主演:(.+)?类型:(.+)?地区:(.+)?语言:(.+)?字幕:(.+)?上映日期:(.+)?剧情简介(.+)`)

var coverImg = "#post_content>p>img"
var info = "#post_content>p"

func FilmDetail(content []byte, fatherUrl string) engine.ParseResult{
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
	if len(infoImg) < 1 {
		log.Printf("详情图未找到： %s", fatherUrl)
	}
	parse.Requests = requests
	parse.Items = items
	return engine.ParseResult{}
}

func FilmDetailByDuc(content []byte, title string) engine.ParseResult{

	cont := bytes.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(cont)

	var parse engine.ParseResult
	//var requests []engine.Request
	//var items []interface{}

	item := model.Film{
		Name: title,
		Type: 3,
		Introduce: model.Introduce{},
	}

	doc.Find(coverImg).Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("src")
		if i == 0 {
			item.Cover = url
		} else {
			item.InfoImage = url
		}
	})

	inf := doc.Find(info).Text()
	result := filmInfo.FindAllSubmatch([]byte(inf), -1)

	for _, v := range result {
		for i := 0; i < len(v); i++ {
			if i == 0 {
				continue
			}

			_, err := film[i]
			if !err {
				continue
			}

			con := string(v[i])
			switch i {
			case 1:
				item.Introduce.Director = con
				break
			case 2:
				item.Introduce.Writer = con
				break
			case 3:
				actors := strings.Split(con, " / ")
				item.Introduce.Actors = actors
				break
			case 4:
				break
			case 5:
				break
			case 6:
				break
			case 9:
				item.Introduce.Plot = con
				break
			}
		}
	}

	log.Printf("%v", item)
	return parse
}
