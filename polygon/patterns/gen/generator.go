package gen

import "fmt"

func GeneratorExample() {
	for v := range Gen(1, 5) {
		fmt.Println(v)
	}
}

func Gen(start, end int) <-chan int {
	result := make(chan int)
	go func() {
		for i := start; i < end; i++ {
			result <- i
		}
		close(result)
	}()

	return result
}
