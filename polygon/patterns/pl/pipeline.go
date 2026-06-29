package pl

import (
	"fmt"
	"time"
)

func PipelineExample() {
	stream := make(chan int, 5)
	for i := 0; i < 5; i++ {
		stream <- i
	}

	close(stream)

	done := make(chan struct{})
	defer close(done)
	start := time.Now()
	pipeline := Add(done, Add(done, stream, 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}

	fmt.Println("worked is - ", time.Since(start))
}

func Add(done <-chan struct{}, stream <-chan int, delta int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for v := range stream {
			time.Sleep(time.Microsecond * 3)
			select {
			case <-done:
				return
			case outCh <- v + delta:

			}
		}
	}()

	return outCh
}
