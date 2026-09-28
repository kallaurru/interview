package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var stor atomic.Uint64

	countWorkers := 800
	countEvents := 2000
	wg := sync.WaitGroup{}
	wg.Add(countWorkers)
	for i := 0; i < countWorkers; i++ {
		go worker(&wg, &stor, countEvents)
	}

	wg.Wait()

	fmt.Printf("%d\n", stor.Load())
}

func worker(wg *sync.WaitGroup, stor *atomic.Uint64, events int) {
	defer wg.Done()
	const cost uint64 = 7

	for i := 0; i < events; i++ {
		stor.Add(cost)
	}
}
