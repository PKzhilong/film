package schedules

import "filmspider/engine"

type SimpleSchedule struct {
	workerChannel chan engine.Request
}

func (s *SimpleSchedule) Submit(request engine.Request)  {
	go func() {
		s.workerChannel <- request
	}()
}

func (s *SimpleSchedule) ConfWorkChannel(w chan engine.Request)  {
	s.workerChannel = w
}


