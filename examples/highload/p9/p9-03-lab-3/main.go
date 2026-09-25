package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	chanSize  = 16
	producers = 8
	lim       = 200000
	divider   = 4000
)

func main() {
	var (
		attempts, received atomic.Int64
	)
	resCh := make(chan int, chanSize)
	done := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(producers)
	for i := 0; i < producers; i++ {
		go worker(&wg, &attempts, resCh)
	}
	go func() {
		for range resCh {
			received.Add(1)
		}
		close(done)
	}()

	wg.Wait()
	close(resCh)
	<-done
	fmt.Printf("attempts=%d received=%d\n", attempts.Load(), received.Load())
}

func worker(wg *sync.WaitGroup, counter *atomic.Int64, resCh chan<- int) {
	defer wg.Done()

	for i := 0; i < lim; i++ {
		counter.Add(1)
		if (i+1)%divider == 0 {
			resCh <- 1
		}
	}
}
