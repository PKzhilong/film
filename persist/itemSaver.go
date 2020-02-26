package persist

import (
	"filmspider/engine"
	"filmspider/model"
	"github.com/go-redis/redis/v7"
)

func ItemServer(redis *redis.Client) chan interface{} {
	db := engine.DBRun()
	out := make(chan interface{})

	go func() {
		for {
			item := <-out
			film, ok := item.(model.Film)
			if ok {
				db.Create(&film)
			}
		}
		defer db.Close()
	}()

	return out
}
