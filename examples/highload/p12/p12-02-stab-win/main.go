package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	// Первая строка — длина стабилизационного окна w (секунды)
	if !scanner.Scan() {
		return
	}
	w, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return
	}

	// Вторая строка — количество тактов n
	if !scanner.Scan() {
		return
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return
	}

	type tick struct {
		t   int
		rec int
	}

	// Читаем все такты
	ticks := make([]tick, 0, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}
		t, _ := strconv.Atoi(fields[0])
		rec, _ := strconv.Atoi(fields[1])
		ticks = append(ticks, tick{t, rec})
	}

	deque := make([]int, 0, len(ticks))
	results := make([]int, 0, len(ticks))

	for i, tk := range ticks {
		// Убираем из начала элементы, вышедшие за левую границу окна
		for len(deque) > 0 && ticks[deque[0]].t < tk.t-w {
			deque = deque[1:]
		}
		// Поддерживаем убывание рекомендаций: удаляем с конца всё, что <= текущей
		for len(deque) > 0 && ticks[deque[len(deque)-1]].rec <= tk.rec {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		// Максимум в окне — в голове дека
		results = append(results, ticks[deque[0]].rec)
	}

	for i, v := range results {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
