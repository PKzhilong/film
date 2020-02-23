package main

import (
	"filmspider/6v/parse"
	"filmspider/engine"
	"filmspider/schedules"
)

func main()  {
	//eg := engine.SimpleEngine{}
	eg := engine.CronEngine{
		Scheduler: &schedules.QueueSchedule{},
		WorkerChannelCount: 10,
	}
	eg.Run(engine.Request{
		//Url: "https://www.i6v.cc",
		//Url: "https://www.i6v.cc/xijupian/866.html",
		Url: "https://www.i6v.cc/zhanzhengpian/12849.html",
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return parse.FilmDetailByDuc(bytes, "2323")
		},
		//ParserFunc: parse.CategoryParse,
	})
}
