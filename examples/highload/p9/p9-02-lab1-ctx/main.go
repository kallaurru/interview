package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const timeout = 40
const limit = 3

func main() {
	resCh := make(chan int)
	w := 6
	res := 0
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(w)
	for i := 0; i < w; i++ {
		go worker(ctx, &wg, resCh, i)
	}

	for i := 0; i < limit; i++ {
		val, ok := <-resCh
		if !ok {
			break
		}
		res += val
	}
	cancel()
	wg.Wait()
	fmt.Printf("%d\n", res)

}

func worker(ctx context.Context, wg *sync.WaitGroup, resCh chan<- int, id int) {
	defer wg.Done()

	pause := timeout * (id + 1)
	time.Sleep(time.Duration(pause) * time.Millisecond)
	select {
	case <-ctx.Done():
	case resCh <- 1:

	}
}
