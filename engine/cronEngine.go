package engine

import "github.com/jinzhu/gorm"

type CronEngine struct {
	WorkerChannelCount int
	Scheduler Scheduler
	Db *gorm.DB
}

type Scheduler interface {
	Submit(request Request)
	WorkerChannel() chan Request
	Run()
	NotifyWorkChannel
}

type NotifyWorkChannel interface {
	WorkReady(chan Request)
}

func (c *CronEngine) Run(seed ...Request)  {

	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerChannelCount; i++ {
		createWorker(c.Scheduler.WorkerChannel(), out, c.Scheduler)
	}

	//通过调度器收request
	for _, v := range seed  {
		c.Scheduler.Submit(v)
	}

	for {
		result := <- out
		if len(result.Items) > 0 {
			//for _, v := range result.Items {
			//	log.Printf("获取内容： %v", )
			//}
		}

		for _, r := range result.Requests {
			c.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, n NotifyWorkChannel)  {
	go func() {
		for {
			n.WorkReady(in)
			request := <- in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}