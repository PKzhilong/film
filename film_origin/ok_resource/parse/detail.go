package parse

import (
	"bytes"
	"filmspider/engine"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

var aliasName = regexp.MustCompile(`别名.*`)
var director = regexp.MustCompile(`导演.*`)
var actors = regexp.MustCompile(`主演.*`)
var cg = regexp.MustCompile(`类型.*`)
var area = regexp.MustCompile(`地区.*`)
var language = regexp.MustCompile(`语言.*`)
var showTime = regexp.MustCompile(`上映.*`)
var filmLange = regexp.MustCompile(`片长.*`)
var playTimes = regexp.MustCompile(`今日播放量.*`)

func Detail(content []byte, categoryID int, categories *engine.Categories) engine.ParseResult {
	bodyReader := bytes.NewReader(content)
	cont, err := goquery.NewDocumentFromReader(bodyReader)
	var parse engine.ParseResult
	if err != nil {
		return parse
	}
	//...获取基本数据
	getDetailInfo(cont, &parse)
	return parse
}

func getDetailInfo(docu *goquery.Document, parse *engine.ParseResult)  {
	baseInfo := docu.Find(".vodBox")
	cover, _ := baseInfo.Find(".vodImg img").Attr("src")

	info := baseInfo.Find(".vodInfo")
	title := info.Find("h2").Text()
	score := info.Find("label").Text()
	fmt.Printf("标题： %s  得分：%s 封面图： %s\n", title, score, cover)

	var an string
	var dr string
	var at string


	info.Find(".vodinfobox ul li").Each(func(i int, selection *goquery.Selection) {
		//...todo 别名
		if result := aliasName.FindString(selection.Text()); len(result) > 0 {
			an = selection.Find("span").Text()
		}


		fmt.Printf("匹配结果： %v", result)
		//...todo 导演
		if result := director.FindString(selection.Text()); len(result) > 0 {
			dr = selection.Find("span").Text()
		}
		//...todo 主演
		if result := actors.FindString(selection.Text()); len(result) > 0 {
			at = selection.Find("span").Text()
		}
		//...todo 类型
		//...todo 别名
		//...todo 别名

	})



}
