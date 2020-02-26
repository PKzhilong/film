package persist

func ItemServer() chan interface{} {
	out := make(chan interface{})

	go func() {
		for {
			item := <- out

		}

	}()

	return out
}
