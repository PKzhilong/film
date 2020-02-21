package main

import (
	"filmspider/6v/parse"
	"filmspider/engine"
)

func main()  {

	engine.Run(engine.Request{
		Url: "https://www.i6v.cc",
		//Url: "https://www.i6v.cc/xijupian/12510.html",
		//ParserFunc: parse.FilmDetail,
		ParserFunc: parse.CategoryParse,
	})
}
