package wp

import (
	"sync"
	"time"
)

type Task struct {
	ID       int
	WorkerID int
}

func WorkerPoolExample() {
	wg := sync.WaitGroup{}
	result := make(chan Task, 3)
	tasks := make(chan Task, 3)

	for i := 0; i < 3; i++ {
		tasks <- Task{
			ID:       i,
			WorkerID: 0,
		}
	}
	close(tasks)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(i+1, &wg, tasks, result)
	}
	wg.Wait()
	close(result)
}

func worker(id int, wg *sync.WaitGroup, tasks <-chan Task, result chan<- Task) {
	defer wg.Done()
	for task := range tasks {
		time.Sleep(time.Millisecond * 12)
		result <- Task{
			ID:       task.ID,
			WorkerID: id,
		}
	}
}
