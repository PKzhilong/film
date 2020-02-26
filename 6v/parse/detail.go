package parse

import (
	"bytes"
	"filmspider/engine"
	"filmspider/model"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

var coverDetail = regexp.MustCompile(`<p>[^<]*<img alt[=|"]* src="([^"]+)"[^>]*>`)

//导演
var dirtor = regexp.MustCompile(`导　　演([^<]+?)<br/>`)
var dirtor2 = regexp.MustCompile(`导演[\s|:]+?([^<]+?)<br/>`)

//编剧
var writer = regexp.MustCompile(`编　　剧([^<]+?)<br/>`)
var writer2 = regexp.MustCompile(`编剧[\s|:]+?([^<]+?)<br/>`)

//主演
var actor = regexp.MustCompile(`主　　演([^<]+?)<br/>`)
var actor2 = regexp.MustCompile(`主演[\s|:]+?([^<]+?)<br/>`)
//类型
var filmType = regexp.MustCompile(`类　　别([^<]+?)<br/>`)
var filmType2 = regexp.MustCompile(`类型[\s|:]+?([^<]+?)<br/>`)

//地区
var area = regexp.MustCompile(`国　　家([^<]+?)<br/>`)
var area3 = regexp.MustCompile(`产　　地([^<]+?)<br/>`)
var area2 = regexp.MustCompile(`地区[\s|:]+?([^<]+?)<br/>`)

//字幕
var font = regexp.MustCompile(`字　　幕([^<]+?)<br/>`)
var font2 = regexp.MustCompile(`字幕[\s|:]+?([^<]+?)<br/>`)

//上映时间
var onTime = regexp.MustCompile(`年　　代([^<]+?)<br/>`)
var onTime2 = regexp.MustCompile(`上映日期[\s|:]+?([^<]+?)<br/>`)

var downApp = regexp.MustCompile(`([^：]+).*`)
var downPwd = regexp.MustCompile(`提取码：([\w]+)`)

//下载地址

var coverImg = "#post_content>p>img"
var info = "#post_content"
var downloadTable = "table>tbody tr"
//在线观看地址
var htmlOnline = ".context>.widget.box.row"

func FilmDetailByDuc(content []byte, title string, categoryID int) engine.ParseResult {

	cont := bytes.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(cont)

	var parse engine.ParseResult

	item := model.Film{
		Name: title,
		CategoryID: categoryID,
	}

	doc.Find(coverImg).Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("src")
		if i == 0 {
			item.Cover = url
		}
	})

	inf, _ := doc.Find(info).Html()
	conBytes := []byte(inf)
	item.Director = getContent(conBytes, dirtor, dirtor2)

	item.Writer = getContent(conBytes, writer, writer2)
	item.Actors = getContent(conBytes, actor2, actor)
	item.Language = getContent(conBytes, font2, font)
	item.ShowTime = getContent(conBytes, onTime, onTime2)
	item.Area = getContent(conBytes, area, area2)

	item.TypeName = getContent(conBytes, filmType, filmType2)

	doc.Find(downloadTable).Each(func(i int, selection *goquery.Selection) {
		c := selection.Find("a").Text()
		u := model.DownloadUrl{Title: c}

		url, er := selection.Find("a").Attr("href")
		dt := selection.Find("td").Text()
		//
		dts := downApp.FindAllStringSubmatch(dt, -1)
		dtp := downPwd.FindAllStringSubmatch(dt, -1)

		if er {
			u.Url = url
		}
		if len(dts) > 0 {
			u.DownType = dts[0][1]
		}

		if len(dtp) > 0 {
			u.Password = dtp[0][1]
		}

		item.DownloadUrls = append(item.DownloadUrls, u)
	})

	//获取在线观看详情页
	doc.Find(htmlOnline).Each(func(i int, selection *goquery.Selection) {
		var sh model.HtmlOnline

		url, err := selection.Find("a").Attr("href")
		if err {
			sh.ParentUrl = url
		}
		name := selection.Find("a").Text()
		sh.Name = name
		item.HtmlOnlines = append(item.HtmlOnlines, sh)
	})
	//parse.Requests = requests
	parse.Items = append(parse.Items, item)
	return parse
}

func getContent(cont []byte, regs ...*regexp.Regexp) string {
	content := ""
	for _, reg := range regs {
		result := reg.FindAllSubmatch(cont, -1)
		if len(result) > 0 {
			content = string(result[0][1])
		}
	}
	return content
}
