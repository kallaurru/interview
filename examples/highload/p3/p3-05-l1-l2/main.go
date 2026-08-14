package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type LRUCache struct {
	router   map[string]*list.Element
	items    *list.List
	capacity int
	stat     Stat
	logOn    bool
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		router:   make(map[string]*list.Element, capacity+1),
		items:    list.New(),
		capacity: capacity,
		logOn:    false,
		stat:     Stat{},
	}
}

func (c *LRUCache) Get(key string) bool {
	node, ok := c.router[key]
	if !ok {
		c.log(false)
		return false
	}
	c.stat.L1Hits++
	c.items.MoveToFront(node)
	c.log(true)

	return true
}

func (c *LRUCache) Set(key string) {
	node, ok := c.router[key]
	if !ok {
		c.insert(key)
		return
	}
	c.items.MoveToFront(node)
}

func (c *LRUCache) Stat() Stat {
	return c.stat
}

func (c *LRUCache) insert(key string) {
	c.capacity--
	node := c.items.PushFront(key)
	c.router[key] = node
	if c.capacity < 0 {
		c.evict()
	}
}

func (c *LRUCache) evict() {
	c.capacity++
	node := c.items.Back()
	key, ok := node.Value.(string)
	if !ok {
		return
	}
	delete(c.router, key)
	c.items.Remove(node)
}

func (c *LRUCache) log(isHit bool) {
	if !c.logOn {
		return
	}
	// отключаем в этой задаче log
	if isHit {
		fmt.Printf("%s\n", "HIT")
		return
	}
	fmt.Printf("%s\n", "MISS")
}

type Stat struct {
	L1Hits  int
	L2Hits  int
	DBReads int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var c, n int

	_, err := fmt.Fscan(r, &c)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	cache := New(c)
	cacheL2 := make(map[string]struct{}, n)
	stat := Stat{}

	for i := 0; i < n; i++ {
		// разбираем линию на запчасти
		var key = ""
		_, err = fmt.Fscan(r, &key)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}

		ok := cache.Get(key)
		if ok {
			continue // в L1 кэше счетчик Hits
		}

		_, ok = cacheL2[key]
		if ok {
			stat.L2Hits++
			cache.Set(key)
			continue
		}
		// в обоих кэшах пусто
		stat.DBReads++
		cache.Set(key)
		cacheL2[key] = struct{}{}
	}
	statL1 := cache.Stat()

	fmt.Printf("l1_hits=%d l2_hits=%d db_reads=%d\n",
		statL1.L1Hits, stat.L2Hits, stat.DBReads)
}
