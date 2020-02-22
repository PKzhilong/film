package engine

type CronEngine struct {
	WorkerChannelCount int
	Scheduler Scheduler
}

type Scheduler interface {
	Submit(request Request)
	ConfWorkChannel(chan Request)
}

func (c *CronEngine) Run(seed ...Request)  {

	in := make(chan Request)
	out := make(chan ParseResult)

	c.Scheduler.ConfWorkChannel(in)

	for i := 0; i < c.WorkerChannelCount; i++ {
		createWorker(in, out)
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

func createWorker(in chan Request, out chan ParseResult)  {
	go func() {
		for {
			request := <- in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}