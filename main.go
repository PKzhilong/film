package main

import (
	"filmspider/6v/parse"
	"filmspider/engine"
	"filmspider/model"
	"filmspider/persist"
	"filmspider/repository"
	"filmspider/schedules"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"path"
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

	//sc := &schedules.SimpleSchedule{}
	sc := &schedules.QueueSchedule{}
	sc.Run()

	eg := engine.CronEngine{
		Scheduler: sc,
		WorkerChannelCount: 5,
		ItemChan: itemChan,
		Redis: redisClient,
	}

	//...todo 简单调度器的时候
	//go RunOn(sc.WorkerChannel())

	eg.Run(engine.Request{
		//Url: "https://www.i6v.cc",
		//Url: "https://www.i6v.cc/xijupian/866.html",
		Url: "https://www.i6v.cc/zhanzhengpian/12849.html",

		//Url:  "https://www.i6v.cc/donghuapian/10265.html",
		//Url:  "https://www.i6v.cc/dongzuopian/12791.html",
		//Url: "https://www.i6v.cc/donghuapian/index.html",
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return parse.FilmDetailByDuc(bytes, "2323", 2)
		},
		//ParserFunc: parse.CategoryDetail,
		//ParserFunc: parse.CategoryParse,
	})

}

// 在线地址存储
func RunOn(in chan engine.Request)  {
	db := engine.DBRun()

	count := 0
	db.Model(&model.HtmlOnline{}).Where("play_url = ''").Count(&count)
	limit := 10
	offset := 0

	for offset <= count {
		var list []model.HtmlOnline
		db.Where("play_url = ''").Offset(offset).Limit(limit).Find(&list)
		offset = offset + limit
		for _, v := range list {
			//...todo 如果更新
			basename := path.Base(v.ParentUrl)
			ext := path.Ext(basename)
			//fmt.Printf("获取到的地址：%s 扩展名: %s\n", basename, ext)
			if ext == ".mp4" {
				v.PlayUrl = v.ParentUrl
				repository.HtmlOnline{DB: db}.Update(&v)
				continue
			}


			ID := v.ID
			request := engine.Request{
				Url: v.ParentUrl,
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return parse.OnlineHtml(bytes, ID)
				},
			}
			in <- request
		}
	}
}