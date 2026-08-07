package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	const limit = 3
	var (
		n int
	)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(1)
	}

	if n <= 0 {
		os.Exit(11)
	}

	grR := make(map[string]int, n)
	grS := make(map[string]int64, n)
	paths := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var (
			route string
			count int64
		)
		_, err = fmt.Fscan(r, &route, &count)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		grR[route] += 1
		grS[route] += count
	}

	for path, _ := range grR {
		paths = append(paths, path)
	}
	sort.Slice(paths, func(i, j int) bool {
		if grS[paths[i]] != grS[paths[j]] {
			return grS[paths[i]] > grS[paths[j]] // сумма по убыванию
		}
		return paths[i] < paths[j] // при равенстве — path по возрастанию
	})
	mx := limit
	if len(paths) < mx {
		mx = len(paths)
	}
	for i := 0; i < mx; i++ {
		path := paths[i]
		count := grR[path]
		total := grS[path]
		fmt.Printf("%s %d %d\n", path, total, total/int64(count))
	}
}
