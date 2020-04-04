package run

import (
	"filmspider/engine"
	"filmspider/film_origin/ok_resource/parse"
)

func Run(eg *engine.CronEngine)  {
	eg.Run(engine.Request{
		//Url: "http://www.okzyw.com/",
		Url: "http://www.okzyw.com/?m=vod-detail-id-45982.html",
		//ParserFunc: func(bytes []byte) engine.ParseResult {
		//	return parse.Category(bytes, eg.DB, eg.Categories)
		//},
		ParserFunc: func(bytes []byte) engine.ParseResult {

			return parse.Detail(bytes, 2, eg.Categories)
		},
	})
}