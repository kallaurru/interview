package sema

import (
	"fmt"
	"time"
)

type Semaphore struct {
	sig chan struct{}
}

func New(lim int) *Semaphore {
	return &Semaphore{
		sig: make(chan struct{}, lim),
	}
}

func (s *Semaphore) Acquire() {
	s.sig <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.sig
}
func SemaphoreExample() {
	s := New(2)
	for i := 0; i < 10; i++ {
		s.Acquire()
		go func(s *Semaphore, i int) {
			defer s.Release()
			time.Sleep(time.Millisecond * 12)
			fmt.Println("i -", i)
		}(s, i)
	}
}
