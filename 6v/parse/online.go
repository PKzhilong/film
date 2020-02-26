package parse

import (
	"filmspider/engine"
	"log"
	"regexp"
)

var onlineUrl = regexp.MustCompile(`<iframe .*?src="([^"]+)"></iframe>`)

func OnlineHtml(content []byte, requestID int) engine.ParseResult {
	var parse engine.ParseResult
	//var requests []engine.Request
	//fmt.Printf("%s requestId: %d", content, requestID)
	result := onlineUrl.FindAllSubmatch(content, -1)
	if len(result) < 1 {
		return parse
	}
	log.Printf("观看地址： %s", string(result[0][1]))

	return parse
}