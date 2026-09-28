package main

import (
	"fmt"
	"sync"
)

func main() {
	stor := make(map[string]int, 4)
	countWorkers := 50
	incr := 8000
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	for _, tag := range []string{"cpu", "mem", "net", "disk"} {
		wg.Add(countWorkers)
		for j := 0; j < countWorkers; j++ {
			go worker(&wg, &mx, stor, tag, incr)
		}
	}

	wg.Wait()
	counter := 0
	for _, val := range stor {
		counter += val
	}

	fmt.Printf("%d\n", counter)
}

func worker(wg *sync.WaitGroup, mx *sync.Mutex, stor map[string]int, tag string, events int) {
	defer wg.Done()

	for i := 0; i < events; i++ {
		mx.Lock()
		stor[tag] += 1
		mx.Unlock()
	}
}
