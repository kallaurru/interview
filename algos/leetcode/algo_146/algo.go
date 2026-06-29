package algo_146

import "container/list"

type LRUCache struct {
	cap  int
	len  int
	stor map[int]*list.Element
	used *list.List
}

const constLimitCapacity = 3000
const constNotFoundValue = -1

func New(capacity int) *LRUCache {
	if capacity > constLimitCapacity {
		capacity = constLimitCapacity
	}

	return &LRUCache{
		cap:  capacity,
		len:  0,
		stor: make(map[int]*list.Element, capacity),
		used: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	val, ok := c.stor[key]
	if !ok {
		return constNotFoundValue
	}
	c.moveToHead(val)
	return val.Value.(int)
}

func (c *LRUCache) Put(key, value int) {
	val, ok := c.stor[key]
	if ok {
		val.Value = value
		return
	}
	// нужно вставить
	if c.len == c.cap {
		c.used.Remove(c.used.Back()) // удаляем	последний
		c.len -= 1
	}
	el := c.used.PushFront(value)
	c.stor[key] = el
	c.len += 1
}

func (c *LRUCache) Cap() int {
	return len(c.stor)
}

func (c *LRUCache) moveToHead(e *list.Element) {
	if c.used.Front() == e {
		return
	}
	c.used.MoveToFront(e)
}
