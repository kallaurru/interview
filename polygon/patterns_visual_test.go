package polygon

import (
	"fmt"
	"github.com/kallaurru/interview/polygon/patterns/ffio"
	"github.com/kallaurru/interview/polygon/patterns/gen"
	"github.com/kallaurru/interview/polygon/patterns/pf"
	"github.com/kallaurru/interview/polygon/patterns/pl"
	"github.com/kallaurru/interview/polygon/patterns/sema"
	"github.com/kallaurru/interview/polygon/patterns/wp"
	"testing"
	"time"
)

func TestPF(t *testing.T) {
	pf.PromiseFutureExample()
}

func TestSemaphore(t *testing.T) {
	sema.SemaphoreExample()
}

func TestGenerator(t *testing.T) {
	gen.GeneratorExample()
}

func TestPipeline(t *testing.T) {
	pl.PipelineExample()
}

func TestWorkerPool(t *testing.T) {
	wp.WorkerPoolExample()
}

func TestFanInFunOut(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	stream := make(chan int, 5)
	fanOut := make([]<-chan int, 0, 5)
	for i := 0; i < 5; i++ {
		stream <- i
		fanOut = append(fanOut, pl.Add(done, stream, 1))
	}
	close(stream)

	start := time.Now()
	pipeline := pl.Add(done, ffio.FanIn(done, fanOut...), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
	fmt.Println("worked is -", time.Since(start))
}
