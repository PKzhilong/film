package run

import (
	"filmspider/engine"
	"filmspider/gadfilm/parse"
)

func Run(eg *engine.CronEngine)  {
	eg.Run(engine.Request{
		Url: "http://www.btbtdy.la/screen/0-----time-1.html",
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return parse.Categories(bytes, eg.DB)
		},
		//Url: "http://www.btbtdy.la/btdy/dy19016.html",
		//ParserFunc: func(c []byte) engine.ParseResult {
		//	return parse.Detail(c, eg.DB, 9694, 111)
		//},
		//Url: "http://www.btbtdy.la/vidlist/19016.html",
		//ParserFunc: func(bytes []byte) engine.ParseResult {
		//	return parse.DetailHidden(bytes, eg.DB, 9694)
		//},
		//Url: "http://www.btbtdy.la/play/19016-0-65.html",
		//ParserFunc: func(c []byte) engine.ParseResult {
		//	return parse.OnlineList(c, eg.DB, 8551)
		//},
	})
}

