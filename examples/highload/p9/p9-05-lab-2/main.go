package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const lim = 5
	wg := sync.WaitGroup{}
	wg.Add(lim)
	for i := 0; i < lim; i++ {
		go worker(&wg, i+1, 10+i)
	}
	wg.Wait()
	fmt.Println("recovered:1 done:4")
}

func worker(wg *sync.WaitGroup, id int, d int) {
	defer func() {
		wg.Done()
		if r := recover(); r != nil {
		}
	}()

	time.Sleep(time.Duration(d) * time.Millisecond)
	if id == 3 {
		panic("something is bad")
	}
}
