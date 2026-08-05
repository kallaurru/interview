package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type GGroup struct {
	Count int
	State string
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, t int
	delim := ";"

	_, err := fmt.Fscan(r, &t)
	if err != nil {
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(2)
	}

	if n <= 0 {
		fmt.Printf("%s\n", "NO LEAKS")
		os.Exit(0)
	}

	data := make([]GGroup, 0, n/2)
	states := make(map[string]int, n)

	for i := 0; i < n; i++ {
		var name string
		var state string

		_, err = fmt.Fscan(r, &state, &name)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		state = fmt.Sprintf("%s%s%s", state, delim, name)
		states[state] += 1
	}
	// добавляем группы
	for state, count := range states {
		if count >= t {
			data = append(data, GGroup{State: state, Count: count})
		}
	}
	if len(data) == 0 {
		fmt.Printf("%s\n", "NO LEAKS")
		os.Exit(0)
	}
	sort.Slice(data, func(i, j int) bool {
		// Сначала сравниваем Flat по убыванию
		if data[i].Count != data[j].Count {
			return data[i].Count > data[j].Count // true — i идёт раньше j
		}
		// При равных State — по Name лексикографически (по возрастанию)
		return data[i].State < data[j].State
	})
	for _, item := range data {
		fields := strings.Split(item.State, delim)
		if len(fields) < 2 {
			continue
		}
		fmt.Printf("%d %s %s\n", item.Count, fields[0], fields[1])
	}
}
