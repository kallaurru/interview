package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Free []int

func (f Free) Len() int           { return len(f) }
func (f Free) Less(i, j int) bool { return f[i] < f[j] }
func (f Free) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func (f *Free) Push(x interface{}) {
	*f = append(*f, x.(int))
}

func (f *Free) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}

func (f *Free) Peek() int {
	if f.Len() == 0 {
		return -1
	}
	return (*f)[0]
}
func (f *Free) Update(i, newValue int) bool {
	if i < 0 || i >= len(*f) {
		return false
	}
	(*f)[i] = newValue
	heap.Fix(f, i)
	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var c, t, n int

	_, err := fmt.Fscan(r, &c, &t, &n)
	if err != nil {
		os.Exit(1)
	}

	f := &Free{}
	makeSpan, maxWait := 0, 0

	heap.Init(f)
	for i := 0; i < c; i++ {
		heap.Push(f, 0)
	}
	for i := 0; i < n; i++ {
		a := 0
		_, err = fmt.Fscan(r, &a)
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		start := a
		if f.Peek() > start {
			start = f.Peek()
		}
		if maxWait < start-a {
			maxWait = start - a
		}
		span := start + t
		if makeSpan < span {
			makeSpan = span
		}
		f.Update(0, span)
	}

	fmt.Printf("makespan=%d\n", makeSpan)
	fmt.Printf("max_wait=%d\n", maxWait)
}
