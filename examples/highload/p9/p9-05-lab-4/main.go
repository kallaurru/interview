package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	t := time.Now()
	go func() {
		lim := 4
		eCh := make(chan struct{}, lim)
		wg := sync.WaitGroup{}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		wg.Add(lim)
		for i := 0; i < lim; i++ {
			if i > 0 {
				go worker(ctx, &wg, eCh, 3000*time.Millisecond)
				continue
			}
			go worker(ctx, &wg, eCh, 100*time.Millisecond)
		}
		go func() {
			select {
			case <-eCh:
				cancel()
				return
			case <-ctx.Done():
				return
			}
		}()
		wg.Wait()
		close(eCh)
		close(done)
	}()

	<-done
	diff := time.Since(t)
	fmt.Printf("%d\n", diff.Milliseconds())
}

func worker(ctx context.Context, wg *sync.WaitGroup, eCh chan<- struct{}, d time.Duration) {
	ticker := time.NewTicker(d)
	defer wg.Done()
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return
	case <-ticker.C:
		eCh <- struct{}{}
		return
	}
}
