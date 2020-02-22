package main

import (
	"filmspider/6v/parse"
	"filmspider/engine"
	"filmspider/schedules"
)

func main()  {

	//eg := engine.SimpleEngine{}
	eg := engine.CronEngine{
		Scheduler: &schedules.SimpleSchedule{},
		WorkerChannelCount: 10,
	}
	eg.Run(engine.Request{
		Url: "https://www.i6v.cc",
		//Url: "https://www.i6v.cc/juqingpian/12772.html",
		//ParserFunc: func(bytes []byte) engine.ParseResult {
		//	return parse.FilmDetailByDuc(bytes, "2323")
		//},
		ParserFunc: parse.CategoryParse,
	})
}
