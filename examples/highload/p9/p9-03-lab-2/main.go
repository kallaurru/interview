package main

import (
	"fmt"
	"sync"
)

func main() {
	ch0 := worker(0)
	ch1 := worker(1)
	ch2 := worker(2)
	res := 0
	out := funIn(ch0, ch1, ch2)

	for val := range out {
		res += val
	}
	fmt.Printf("%d\n", res)
}

func worker(id int) chan int {
	out := make(chan int)
	go func(out chan<- int) {
		for j := 1; j <= 10; j++ {
			out <- id*100 + j
		}
		close(out)
	}(out)
	return out
}

func funIn(input ...chan int) chan int {
	out := make(chan int, len(input))
	if len(input) == 0 {
		close(out)
		return out
	}
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(input))
		for _, ch := range input {
			go func(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
				defer wg.Done()
				for val := range in {
					out <- val
				}
			}(&wg, ch, out)
		}
		go func(wg *sync.WaitGroup, out chan int) {
			wg.Wait()
			close(out)
		}(&wg, out)
	}()
	return out
}
