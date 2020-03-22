package parse

import (
	"filmspider/engine"
	"filmspider/model"
	"log"
	"regexp"
)

var onlineUrl = regexp.MustCompile(`<iframe .*?src="([^"]+)"></iframe>`)

func OnlineHtml(content []byte, onlineID int) engine.ParseResult {
	var parse engine.ParseResult
	result := onlineUrl.FindAllSubmatch(content, -1)
	if len(result) < 1 {
		return parse
	}

	playUrl := string(result[0][1])
	parse.Items = append(parse.Items, model.HtmlOnline{
		ID: onlineID,
		PlayUrl: playUrl,
	})
	log.Printf("观看地址： %s", string(result[0][1]))
	return parse
}
