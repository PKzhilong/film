package parse

import (
	"bytes"
	"filmspider/engine"
	"filmspider/model"
	"filmspider/repository"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
)

//获取隐藏地址
var hiddenRequestUrl = "%s/vidlist/%d.html"

func Detail(c []byte, d *gorm.DB, filmID int, urlID int) engine.ParseResult  {
	db = d
	bodyReader := bytes.NewReader(c)
	cont, err := goquery.NewDocumentFromReader(bodyReader)

	var parse engine.ParseResult
	if err != nil {
		fmt.Printf("%v", err)
		return parse
	}
	//fmt.Printf("%s", c)

	film := model.Film{
		ID: filmID,
	}

	cont.Find(".vod .vod_intro.rt").Each(func(i int, selection *goquery.Selection) {

		selection.Find("dl dt").Each(func(index int, s *goquery.Selection) {
			//fmt.Printf("%s\n", s.Text())

			if s.Text() == "语言:" {
				dd := s.Next()
				film.Language = dd.Text()
			}

			if s.Text() == "主演:" {
				dd := s.Next()
				film.Actors = dd.Text()
			}
		})

		film.Plot = selection.Find(".des span").Text()

	})

	er := repository.Film{DB:db}.UpdateByID(filmID, &film)
	if er != nil {
		fmt.Printf("更新电影失败：%d %v", filmID, err)
	}
	hiddenUrl := fmt.Sprintf(hiddenRequestUrl, host, urlID)
	parse.Requests = append(parse.Requests, engine.Request{
		Url: hiddenUrl,
		ParserFunc: func(c []byte) engine.ParseResult {
			return DetailHidden(c, db, filmID)
		},
	})
	return parse
}