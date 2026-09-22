package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const limWorkers = 5

func main() {
	wg := sync.WaitGroup{}
	imitations := []int{50, 100, 150, 300, 400}
	resCh := make(chan int, limWorkers)
	result := 0
	wg.Add(len(imitations))
	for _, imit := range imitations {
		go worker(&wg, imit, resCh)
	}
	wg.Wait()
	close(resCh)

	for i := 0; i < limWorkers; i++ {
		result += <-resCh
	}

	fmt.Println(result)
}

func worker(wg *sync.WaitGroup, imitation int, resCh chan<- int) {
	t := time.NewTicker(10 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)

	defer wg.Done()
	defer t.Stop()
	defer cancel()

	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			counter += 10
			if ctx.Err() != nil {
				return
			}
			if counter >= imitation {
				resCh <- 1
				return
			}
		}
	}
}
