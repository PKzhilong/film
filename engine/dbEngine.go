package engine

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

func DBRun() *gorm.DB {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	sqlConf := fmt.Sprintf("%s:%v@(%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))
	db, err := gorm.Open("mysql", sqlConf)
	if err != nil{
		panic(err)
	}

	return db
}