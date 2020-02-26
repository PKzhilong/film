package persist

import (
	"filmspider/engine"
	"filmspider/model"
)

func ItemServer() chan interface{} {
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
