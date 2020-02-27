package main

import (
	"filmspider/6v/parse"
	"filmspider/engine"
	"filmspider/model"
	"filmspider/persist"
	"filmspider/schedules"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main()  {

	err := godotenv.Load(".env")
	if err != nil {
		panic("获取环境变量失败")
	}

	redisClient := engine.RedisRun()
	db := engine.DBRun()
	itemChan := persist.ItemServer(redisClient, db)

	//eg := engine.SimpleEngine{}
	schedule := &schedules.SimpleSchedule{}
	schedule.Run()

	eg := engine.CronEngine{
		Scheduler: schedule,
		WorkerChannelCount: 5,
		ItemChan: itemChan,
		Redis: redisClient,
	}
	go RunOn(schedule.WorkerChannel())


	eg.Run(engine.Request{
		//Url: "https://www.i6v.cc",
		//Url: "https://www.i6v.cc/xijupian/866.html",
		//Url: "https://www.i6v.cc/zhanzhengpian/12849.html",

		Url:  "https://www.i6v.cc/zhanzhengpian/12849.html",
		//Url: "https://www.i6v.cc/donghuapian/index.html",
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return parse.FilmDetailByDuc(bytes, "2323", 2)
		},
		//ParserFunc: parse.CategoryDetail,
		//ParserFunc: parse.CategoryParse,
	})

}
