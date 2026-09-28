package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct {
	Limit    int
	Checksum int
}

func New(lim int) *Config {
	return &Config{Limit: lim, Checksum: lim * 2}
}

func (c *Config) Update(lim int) {
	tmpLim := c.Limit + lim
	c.Limit = tmpLim
	c.Checksum = tmpLim * 2
}

func main() {
	var safeCfg atomic.Pointer[Config]
	var badCfg atomic.Int64

	countReaders := 20
	steps := 1000
	wg := sync.WaitGroup{}
	config := New(1)
	safeCfg.Store(config)
	wg.Add(1)

	go func(wg *sync.WaitGroup, sCfg atomic.Pointer[Config]) {
		const lim = 1000

		defer wg.Done()
		for i := 1; i < lim; i++ {
			old := sCfg.Load()
			old.Update(1)
			sCfg.Store(old)
		}
	}(&wg, safeCfg)

	wg.Add(countReaders)
	for i := 0; i < countReaders; i++ {
		go worker(&wg, safeCfg, &badCfg, steps)
	}
	wg.Wait()
	fmt.Printf("%d\n", badCfg.Load())
}

func worker(wg *sync.WaitGroup, sCfg atomic.Pointer[Config], counter *atomic.Int64, steps int) {
	defer wg.Done()

	for i := 0; i < steps; i++ {
		cfg := sCfg.Load()
		if cfg.Checksum/cfg.Limit == 2 {
			continue
		}
		counter.Add(1)
	}
}
