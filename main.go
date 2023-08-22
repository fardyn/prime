package main

func generator(limit int, ch chan<- int) {

	for i := 2; i < limit; i++ {
		ch <- i
	}

	close(ch)

}
