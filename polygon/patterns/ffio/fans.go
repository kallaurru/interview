package ffio

import "sync"

func FanIn(done <-chan struct{}, items ...<-chan int) <-chan int {
	result := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range items {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			select {
			case <-done:
				return
			case v := <-ch:
				result <- v
			}
		}(ch)
	}
	go func() {
		defer close(result)
		wg.Wait()
	}()

	return result
}
