package run

import (
	"filmspider/engine"
	"filmspider/film_origin/ok_resource/parse"
)

func Run(eg *engine.CronEngine)  {
	eg.Run(engine.Request{
		Url: "http://www.okzyw.com/",
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return parse.Category(bytes, eg.DB)
		},

	})
}