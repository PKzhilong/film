package parse

import (
	"bytes"
	"filmspider/engine"
	"filmspider/model"
	"filmspider/repository"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var detailID = regexp.MustCompile(`dy([0-9]+)\.html`)

func CategoryDetail(c []byte, d *gorm.DB, categoryID int, parentUrl string) engine.ParseResult {
	bodyReader := bytes.NewReader(c)
	content, err := goquery.NewDocumentFromReader(bodyReader)

	var parse engine.ParseResult
	if err != nil {
		return parse
	}

	content.Find(".list_su ul>li").Each(func(i int, selection *goquery.Selection) {
		show := selection.Find(".pic_link")
		imgUrl, _ := show.Find("img").Attr("data-src")

		name, _ := show.Attr("title")
		url, _ := show.Attr("href")

		desc := selection.Find(".cts_ms")
		score := desc.Find(".title span").Text()
		categories := desc.Find(".des").Text()
		attrList := handleAttr(categories)


		film := model.Film{
			Name: name,
			Cover: imgUrl,
			Type: categoryID,
			Area: attrList["area"],
			ShowTime: attrList["year"],
			TypeName: attrList["categories"],
		}
		repository.Film{DB: db}.Create(&film)
		if film.ID == 0 {
			fmt.Printf("错误信息： name: %s url: %s img: %s score: %d cl: %s film: %v\n", name, url, imgUrl, score, attrList, film)
			return
		}

		requestUrl := host + url
		urlID := getUrlID(requestUrl)
		request := engine.Request{
			Url: requestUrl,
			ParserFunc: func(c []byte) engine.ParseResult {
				return Detail(c, db, film.ID, urlID)
			},
		}
		parse.Requests = append(parse.Requests, request)
		fmt.Printf("name: %s url: %s img: %s score: %d cl: %s\n", name, url, imgUrl, score, attrList)
	})

	//...todou 获取下一页
	content.Find(".pages a").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text != "下一页" {
			return
		}
		u, _ := selection.Attr("href")
		ru := getNextUrl(parentUrl, u)
		request := engine.Request{
			Url: ru,
			ParserFunc: func(i []byte) engine.ParseResult {
				return CategoryDetail(i, db, categoryID, parentUrl)
			},
		}
		parse.Requests = append(parse.Requests, request)
	})


	return parse
}

func handleAttr(categories string) map[string]string {
	attrs := make(map[string]string)
	if len(categories) < 1 {

	}

	list := strings.Split(categories, " ")

	if len(list) == 3 {

		attrs["year"] = strings.Replace(list[0], "年", "", -1)
		attrs["area"] = list[1]
		attrs["categories"] = list[2]
	}
	return attrs
}

func getUrlID(url string) int {
	result := detailID.FindAllStringSubmatch(url, -1)
	if len(result) < 1 {
		fmt.Printf("获取UrlID失败： %s", url)
		return 0
	}

	id, err := strconv.Atoi(result[0][1])
	if err != nil {
		fmt.Printf("获取UrlID失败： %s", url)
		return 0
	}
	return id
}

func getNextUrl(baseUrl string, newUrl string) string  {
	u, _ := url.Parse(baseUrl)
	p := u.Path

	pl := strings.Split(p,"/")
	pu := "/"
	for i, v := range pl {
		if len(pl) == (i+1) {
			pu += "/"
			continue
		}
		pu += v
	}

	resultUrl := host + pu + newUrl
	return resultUrl
}
