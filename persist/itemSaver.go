package persist

import (
	"crypto/md5"
	"filmspider/model"
	"filmspider/repository"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

func ItemServer(redis *redis.Client, db *gorm.DB) chan interface{} {

	out := make(chan interface{})

	go func() {
		for {
			item := <-out
			film, ok := item.(model.Film)
			if ok {
				repository.Film{DB: db}.Create(&film)
			}

			online, ok := item.(model.HtmlOnline)
			if ok {
				repository.HtmlOnline{DB: db}.Update(&online)
			}
		}
	}()

	return out
}

func Md5(name string) string {
	w := []byte(name)
	hex := md5.Sum(w)
	return fmt.Sprintf("%x", hex)
}
