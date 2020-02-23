package schedules

import "filmspider/engine"

type QueueSchedule struct {
	RequestChannel chan engine.Request
	WorkChannel    chan chan engine.Request
}

func (q *QueueSchedule) WorkerChannel() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueSchedule) Submit(request engine.Request) {
	q.RequestChannel <- request
}

func (q *QueueSchedule) WorkReady(r chan engine.Request) {
	q.WorkChannel <- r
}

func (q *QueueSchedule) Run() {
	q.RequestChannel = make(chan engine.Request)
	q.WorkChannel = make(chan chan engine.Request)

	go func() {
		// 维护两个队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request


		for {

			var actionReq engine.Request
			var actionWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				actionReq = requestQ[0]
				actionWorker = workerQ[0]
			}

			select {
			case r := <-q.RequestChannel:
				requestQ = append(requestQ, r)
			case w := <-q.WorkChannel:
				workerQ = append(workerQ, w)
			case actionWorker <- actionReq:
				//if len(requestQ) > 1 {
					requestQ = requestQ[1:]
				//} else {
				//	requestQ = []engine.Request{}
				//}

				//if len(workerQ) > 1 {
					workerQ = workerQ[1:]
				//} else {
				//	workerQ = []chan engine.Request{}
				//}
			}
		}
	}()
}
