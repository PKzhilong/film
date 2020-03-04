package parse

import (
	"bytes"
	"filmspider/engine"
	"filmspider/repository"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"strconv"
)

var host = "http://www.btbtdy.la"
var db *gorm.DB

func Categories(content []byte, d *gorm.DB) engine.ParseResult {
	bodyReader := bytes.NewReader(content)
	cont, err := goquery.NewDocumentFromReader(bodyReader)
	db = d

	var parse engine.ParseResult
	if err != nil {
		return parse
	}

	cont.Find(".list .s_index dl").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("dt").Text()

		switch {
		case title == "选择类型：":
			saveContentType(selection, &parse)
			break
		case title == "选择分类：":
			saveCategory(selection, &parse)
			break
		case title == "选择年代：":
			saveYear(selection, &parse)
			break
		case title == "选择地区：":
			saveArea(selection, &parse)
			break
		case title == "选择语言：":
			saveLanguage(selection, &parse)
			break
		}

	})

	return parse
}

func saveContentType(s *goquery.Selection, p *engine.ParseResult) {
	rep := repository.ContentTypes{DB: db}
	s.Find("dd>a").Each(func(i int, selection *goquery.Selection) {
		value := selection.Text()
		//全部的分类去掉
		if i == 0 {
			return
		}
		url, _ := selection.Attr("href")
		categoryID := rep.CreateIfExist(value)

		ru := host + url
		request := engine.Request{
			Url: ru,
			ParserFunc: func(i []byte) engine.ParseResult {
				return CategoryDetail(i, db, categoryID, ru)
			},
		}
		p.Requests = append(p.Requests, request)
	})
}

func saveCategory(s *goquery.Selection, p *engine.ParseResult) {
	rep := repository.Category{DB: db}
	s.Find("dd>a").Each(func(i int, selection *goquery.Selection) {
		value := selection.Text()
		//全部的分类去掉
		if i == 0 {
			return
		}
		rep.CreateIfNotExist(value)
	})
}

func saveYear(s *goquery.Selection, p *engine.ParseResult) {
	rep := repository.Years{DB: db}
	s.Find("dd>a").Each(func(i int, selection *goquery.Selection) {
		value := selection.Text()
		//全部的分类去掉
		if i == 0 {
			return
		}
		year, _ := strconv.Atoi(value)
		rep.CreateIfNotExist(year)
	})
}

func saveArea(s *goquery.Selection, p *engine.ParseResult) {
	rep := repository.Areas{DB: db}
	s.Find("dd>a").Each(func(i int, selection *goquery.Selection) {
		value := selection.Text()
		//全部的分类去掉
		if i == 0 {
			return
		}
		rep.CreateIfNotExist(value)
	})
}

func saveLanguage(s *goquery.Selection, p *engine.ParseResult) {
	rep := repository.Languages{DB: db}
	s.Find("dd>a").Each(func(i int, selection *goquery.Selection) {
		value := selection.Text()
		//全部的分类去掉
		if i == 0 {
			return
		}
		rep.CreateIfNotExist(value)
	})
}
