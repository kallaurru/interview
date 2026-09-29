package main

import "fmt"

func main() {
	outCh := gen()
	counter := 0
	for val := range outCh {
		counter += val
	}

	fmt.Printf("DONE:%d\n", counter)
}

func gen() chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		const lim = 20
		for i := 0; i < lim; i++ {
			out <- 1
		}
	}()

	return out
}
