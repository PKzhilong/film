package run

import (
	"filmspider/engine"
	"fmt"
)

func Run(eg *engine.CronEngine)  {
	eg.Run(engine.Request{
		Url: "https://91mjw.com/vplay/MjM0NS0xLTE=.html",
		ParserFunc: func(bytes []byte) engine.ParseResult {
			parse := engine.ParseResult{}
			fmt.Printf("%s", bytes)
			return parse
		},

	})
}