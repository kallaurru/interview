package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !sc.Scan() {
		return
	}
	C, err := strconv.Atoi(sc.Text())
	if err != nil || C < 1 {
		return
	}

	cache := make(map[string]*list.Element)
	lru := list.New()

	hits := 0
	misses := 0

	for sc.Scan() {
		code := sc.Text()
		if code == "" {
			continue
		}

		if elem, ok := cache[code]; ok {
			hits++
			lru.MoveToFront(elem)
		} else {
			misses++
			elem := lru.PushFront(code)
			cache[code] = elem
			if lru.Len() > C {
				back := lru.Back()
				if back != nil {
					delete(cache, back.Value.(string))
					lru.Remove(back)
				}
			}
		}
	}

	total := hits + misses
	var rate int
	if total > 0 {
		rate = hits * 100 / total
	} else {
		rate = 0
	}

	fmt.Printf("hits=%d\n", hits)
	fmt.Printf("misses=%d\n", misses)
	fmt.Printf("hit_rate=%d%%\n", rate)
}
