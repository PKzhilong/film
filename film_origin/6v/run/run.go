package run

import (
	"filmspider/film_origin/6v/parse"
	"filmspider/engine"
)

func Run(eg *engine.CronEngine)  {
	eg.Run(engine.Request{
		Url: "https://www.i6v.cc",
		//Url: "https://www.i6v.cc/xijupian/866.html",
		//Url: "https://www.i6v.cc/zhanzhengpian/12849.html",

		//Url:  "https://www.i6v.cc/donghuapian/10265.html",
		//Url:  "https://www.i6v.cc/dongzuopian/12791.html",
		//Url: "https://www.i6v.cc/donghuapian/index.html",
		//ParserFunc: func(bytes []byte) engine.ParseResult {
		//	return parse.FilmDetailByDuc(bytes, "2323", 2)
		//},
		//ParserFunc: parse.CategoryDetail,
		ParserFunc: parse.CategoryParse,
	})
}

