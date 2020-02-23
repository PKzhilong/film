package schedules

import "filmspider/engine"

type SimpleSchedule struct {
	workerChannel chan engine.Request
}

func (s *SimpleSchedule) WorkerChannel() chan engine.Request {
	return s.workerChannel
}

func (s *SimpleSchedule) Submit(request engine.Request)  {
	go func() {
		s.workerChannel <- request
	}()
}


func (s *SimpleSchedule) Run()  {
	s.workerChannel = make(chan engine.Request)
}

func (s *SimpleSchedule) WorkReady(chan engine.Request)  {

}



