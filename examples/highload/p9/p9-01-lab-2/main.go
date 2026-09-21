package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Stats struct {
	current   atomic.Int64
	max       atomic.Int64
	processed atomic.Int64
}

type Pool struct {
	size int
	sema chan struct{}
	wg   sync.WaitGroup
	Stats
}

// New - param ceg - количество одновременно работающих задач
func New(size, ceg int) *Pool {
	return &Pool{size: size, sema: make(chan struct{}, ceg)}
}

func (p *Pool) Close() {
	p.wg.Wait()
}
func (p *Pool) Run() {
	ant := func(wg *sync.WaitGroup, sema chan struct{}) {
		defer func() {
			<-sema
			wg.Done()
			p.Stats.Leave()
		}()
		p.sema <- struct{}{}
		p.Stats.Enter()
		time.Sleep(100 * time.Millisecond)
		p.Stats.Job()
	}

	for i := 0; i < p.size; i++ {
		p.wg.Add(1)
		go ant(&p.wg, p.sema)
	}
}

func (s *Stats) Enter() {
	c := s.current.Add(1)
	for {
		m := s.max.Load()
		if c <= m {
			break
		}
		if s.max.CompareAndSwap(m, c) {
			break
		}
	}
}

func (s *Stats) Leave()           { s.current.Add(-1) }
func (s *Stats) Job()             { s.processed.Add(1) }
func (s *Stats) Current() int64   { return s.current.Load() }
func (s *Stats) Max() int64       { return s.max.Load() }
func (s *Stats) Processed() int64 { return s.processed.Load() }

func main() {
	pool := New(12, 3)
	pool.Run()

	pool.Close()
	fmt.Printf("maxConcurrent=%d processed=%d\n", pool.Stats.Max(), pool.Stats.Processed())
}
