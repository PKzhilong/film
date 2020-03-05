package main

import (
	"filmspider/6v/parse"
	"filmspider/engine"
	"filmspider/gadfilm/run"
	"filmspider/model"
	"filmspider/persist"
	"filmspider/repository"
	"filmspider/schedules"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
	"path"
	"strconv"
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

	wc, _ := strconv.Atoi(os.Getenv("WORKERCOUNT"))
	eg := engine.CronEngine{
		Scheduler: sc,
		WorkerChannelCount: wc,
		ItemChan: itemChan,
		Redis: redisClient,
		DB: db,
	}

	//...todo 简单调度器的时候
	//go RunOn(sc.WorkerChannel())

	//...todo 爬6v
	//run.Run(&eg)


	//...todo 天堂
	run.Run(&eg)
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