package pf

import (
	"fmt"
	"time"
)

func PromiseFutureExample() {
	longRunningTask := func() int {
		time.Sleep(20 * time.Millisecond)
		return 42
	}

	// Запускаем задачу через Promise
	future := Promise(longRunningTask)

	fmt.Println("Задача запущена, можно делать что-то еще...")

	// Ожидаем результат
	result := <-future
	fmt.Println("Результат:", result)
}

func Promise(task func() int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		out <- task()
	}()

	return out
}
