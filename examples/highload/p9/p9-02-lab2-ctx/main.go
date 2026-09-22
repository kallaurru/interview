package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	done := make(chan struct{})
	defer cancel()
	start := time.Now()
	go func(ctx context.Context, done chan struct{}) {
		t := time.NewTicker(10 * time.Millisecond)
		defer t.Stop()
		defer close(done)
		counter := 0
		lim := 2000
		for {
			select {
			case <-ctx.Done():
			case <-t.C:
				counter += 10
				if ctx.Err() != nil {
					return
				}
				if counter >= lim {
					return
				}
			}
		}
	}(ctx, done)
	<-done
	res := time.Since(start)
	if res.Milliseconds() < 300 {
		fmt.Println("OK")
	} else {
		fmt.Println("LATE")
	}
}
