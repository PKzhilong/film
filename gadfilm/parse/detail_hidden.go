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

func DetailHidden(c []byte, d *gorm.DB, filmID int) engine.ParseResult  {

	db = d
	bodyReader := bytes.NewReader(c)
	cont, err := goquery.NewDocumentFromReader(bodyReader)
	var parse engine.ParseResult

	if err != nil {
		fmt.Printf("%v", err)
		return parse
	}

	var htmlOnlines []model.HtmlOnline
	var downloadUrls []model.DownloadUrl

	cont.Find(".p_list").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("h2").Text()
		if title == "云播放" {
			selection.Find("ul li").Each(func(index int, s *goquery.Selection) {
				a := s.Find("a")

				href, _ := a.Attr("href")
				href = host + href
				fmt.Printf("在线地址： %s\n", href)

				online := model.HtmlOnline{
					Name: a.Text(),
					ParentUrl: href,
					FilmID: filmID,
				}
				htmlOnlines = append(htmlOnlines, online)
			})
			return
		}
		selection.Find("ul li").Each(func(index int, s *goquery.Selection) {
			a := s.Find("a")
			title := a.Text()

			link := s.Find("span a")
			downType := link.Text()
			url, _ := link.Attr("href")
			downloadUrl := model.DownloadUrl{
				Title: title,
				DownType: downType,
				Url: url,
				FilmID: filmID,
			}
			downloadUrls = append(downloadUrls, downloadUrl)
		})
	})
	createOnlines(htmlOnlines, &parse)
	createDownloadUrls(downloadUrls)
	return parse
}

func createOnlines(l []model.HtmlOnline, p *engine.ParseResult)  {
	if len(l) < 1 {
		return
	}

	for _, v := range l {
		repository.HtmlOnline{DB: db}.Create(&v)
		if v.ID == 0 {
			continue
		}
		onlineID := v.ID
		request := engine.Request{
			Url: v.ParentUrl,
			ParserFunc: func(c []byte) engine.ParseResult {
				return OnlineList(c, db, onlineID)
			},
		}
		p.Requests = append(p.Requests, request)
	}

}

func createDownloadUrls(l []model.DownloadUrl)  {
	if len(l) < 1 {
		return
	}

	for _, v := range l {
		repository.DownloadUrls{DB: db}.Create(&v)
	}
}