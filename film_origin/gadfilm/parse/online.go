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

func OnlineList(c []byte, d *gorm.DB, onlineID int) engine.ParseResult  {
	db = d
	bodyReader := bytes.NewReader(c)
	cont, err := goquery.NewDocumentFromReader(bodyReader)
	var parse engine.ParseResult

	if err != nil {
		fmt.Printf("%v", err)
		return parse
	}

	iframe := cont.Find("iframe")
	playUrl, _ := iframe.Attr("src")

	htmlOnline := model.HtmlOnline{
		PlayUrl: playUrl,
	}

	fmt.Printf("播放地址(%d)： %s\n", onlineID, playUrl)
	repository.HtmlOnline{DB: d}.UpdateByID(onlineID, &htmlOnline)
	return parse
}