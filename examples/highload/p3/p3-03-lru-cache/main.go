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
}

type Stat struct {
	Hits int
	Miss int
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		router:   make(map[string]*list.Element, capacity+1),
		items:    list.New(),
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key string) {
	node, ok := c.router[key]
	if !ok {
		c.log(false)
		c.stat.Miss++
		return
	}
	c.items.MoveToFront(node)
	c.log(true)
	c.stat.Hits++
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
}

func (c *LRUCache) log(isHit bool) {
	if isHit {
		fmt.Printf("%s\n", "HIT")
		return
	}
	fmt.Printf("%s\n", "MISS")
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var cc, n int

	_, err := fmt.Fscan(r, &cc)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	if n < 1 {
		os.Exit(0) // нет работы у кэша
	}

	cache := New(cc)
	for i := 0; i < n; i++ {
		// разбираем линию на запчасти
		var op, key string

		_, err = fmt.Fscan(r, &op, &key)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		switch op {
		default:
		case "GET":
			cache.Get(key)
		case "PUT":
			cache.Set(key)
		}
	}
	stat := cache.Stat()
	fmt.Printf("hits=%d misses=%d\n", stat.Hits, stat.Miss)
}
