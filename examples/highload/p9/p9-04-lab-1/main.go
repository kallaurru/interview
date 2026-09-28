package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stor := make(map[string]int, 4)
	countWorkers := 200
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	for _, tag := range []string{"cpu", "mem", "net", "disk"} {
		wg.Add(countWorkers)
		for j := 0; j < countWorkers; j++ {
			go worker(&wg, &mx, stor, tag, countWorkers)
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
		time.Sleep(10 * time.Millisecond)
		mx.Lock()
		stor[tag] += 1
		mx.Unlock()
	}
}
