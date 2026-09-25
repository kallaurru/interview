package main

import "fmt"

const (
	constChanSize   = 5
	constItemsCount = 20
)

func main() {
	var (
		sent, dropped, received int
	)
	dataCh := make(chan int, constChanSize)

	for i := 0; i < constItemsCount; i++ {
		select {
		case dataCh <- i:
			sent++
		default:
			dropped++
		}
	}
	close(dataCh)
	for range dataCh {
		received++
	}

	fmt.Printf("sent=%d dropped=%d received=%d\n", sent, dropped, received)
}
