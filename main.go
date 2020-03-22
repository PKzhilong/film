package main

import (
	"filmspider/engine"
	"filmspider/film_origin/ok_resource/run"
	"filmspider/persist"
	"filmspider/repository"
	"filmspider/schedules"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
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
	eg.Categories = initCategories(db)

	//...todo 简单调度器的时候
	//go RunOn(sc.WorkerChannel())

	//...todo 爬6v
	//run.Run(&eg)


	//...todo 天堂
	//run2.Run(&eg)

	//...todo ok网
	run.Run(&eg)
}

func initCategories(db *gorm.DB) *engine.Categories {
	var c engine.Categories
	c.CateList = repository.Category{DB: db}.GetAll()
	lan := &repository.Languages{DB: db}
	c.Languages = lan.GetAll()
	ar := &repository.Areas{DB: db}
	c.Areas = ar.GetAll()

	ye := &repository.Years{DB: db}
	c.Years = ye.GetALL()

	ct := &repository.ContentTypes{DB: db}
	c.ContentTypes = ct.GetAll()
	return &c
}
