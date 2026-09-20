package main

import (
	"fmt"
	"sync"
)

func main() {
	const countW = 5

	chanOut := gen()
	resChan := make(chan int, countW)
	wg := sync.WaitGroup{}
	wg.Add(countW)
	for i := 0; i < countW; i++ {
		go worker(&wg, chanOut, resChan)
	}
	result := 0
	go func(result *int, resCh <-chan int) {
		for res := range resChan {
			*result += res
		}

	}(&result, resChan)
	wg.Wait()
	close(resChan)

	fmt.Println(result)
}

func gen() chan int {
	out := make(chan int)
	go func(out chan<- int) {
		for j := 1; j <= 20; j++ {
			j := j
			out <- j
		}
		close(out)
	}(out)
	return out
}

func worker(wg *sync.WaitGroup, taskCh <-chan int, resCh chan<- int) {
	defer wg.Done()

	for val := range taskCh {
		resCh <- val * 2
	}
}
