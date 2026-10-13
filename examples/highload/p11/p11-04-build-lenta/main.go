package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Post struct {
	id string
	ts int64
}

type Item struct {
	post     Post
	timeline int
	pos      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].post.ts != pq[j].post.ts {
		return pq[i].post.ts > pq[j].post.ts
	}
	return pq[i].post.id < pq[j].post.id
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !sc.Scan() {
		return
	}
	parts := strings.Fields(sc.Text())
	if len(parts) < 2 {
		return
	}
	K, _ := strconv.Atoi(parts[0])
	N, _ := strconv.Atoi(parts[1])

	timelines := make([][]Post, K)

	for i := 0; i < K; i++ {
		if !sc.Scan() {
			return
		}
		m, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
		timelines[i] = make([]Post, 0, m)
		for j := 0; j < m; j++ {
			if !sc.Scan() {
				return
			}
			line := strings.Fields(sc.Text())
			if len(line) < 2 {
				continue
			}
			ts, _ := strconv.ParseInt(line[1], 10, 64)
			timelines[i] = append(timelines[i], Post{id: line[0], ts: ts})
		}
		// Сортируем каждый таймлайн на случай, если входные данные не отсортированы.
		// По условию они уже отсортированы, но это делает решение устойчивым.
		sort.Slice(timelines[i], func(a, b int) bool {
			if timelines[i][a].ts != timelines[i][b].ts {
				return timelines[i][a].ts > timelines[i][b].ts
			}
			return timelines[i][a].id < timelines[i][b].id
		})
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for i, tl := range timelines {
		if len(tl) > 0 {
			heap.Push(pq, &Item{post: tl[0], timeline: i, pos: 0})
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	count := 0
	for pq.Len() > 0 && count < N {
		item := heap.Pop(pq).(*Item)
		fmt.Fprintln(out, item.post.id)
		count++

		nextPos := item.pos + 1
		if nextPos < len(timelines[item.timeline]) {
			heap.Push(pq, &Item{
				post:     timelines[item.timeline][nextPos],
				timeline: item.timeline,
				pos:      nextPos,
			})
		}
	}
}
