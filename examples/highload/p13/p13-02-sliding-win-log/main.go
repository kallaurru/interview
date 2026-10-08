package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RLim struct {
	L, W int
	stor map[string][]int
}

func New(l, w int) *RLim {
	return &RLim{L: l, W: w, stor: make(map[string][]int, 256)}
}

func (rl *RLim) Limiting(key string, ts int) {
	if rl.L == 0 {
		rl.print(false)
		return
	}
	_, ok := rl.lim(key, ts)

	rl.print(ok)
}

func (rl *RLim) lim(key string, ts int) (int, bool) {
	items, ok := rl.stor[key]
	if !ok {
		items = make([]int, 0, 32)
		items = append(items, ts)
		rl.stor[key] = items
		return len(items), true
	}
	start := ts - rl.W
	lim := 0
	for i := 0; i < len(items); i++ {
		if items[i] <= start {
			continue
		}
		lim++
	}
	flag := false
	if lim < rl.L {
		// проходит
		items = append(items, ts)
		rl.stor[key] = items
		flag = true
	} else {
		lim = 0
	}

	return lim, flag
}

func (rl *RLim) print(access bool) {
	if access {
		fmt.Printf("ALLOW\n")
		return
	}
	fmt.Printf("DENY\n")
}

func main() {
	var (
		l, w int
	)
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем первую строку
	if !scanner.Scan() {
		// Если ввода нет, завершаемся или обрабатываем ошибку
		fmt.Println("No input")
		os.Exit(1)
	}
	firstLine := scanner.Text()
	parts := strings.Fields(firstLine) // разделяем по пробелам
	if len(parts) < 2 {
		fmt.Printf("Skipping invalid line: %q\n", 1)
		os.Exit(2)
	}
	l, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(22)
	}

	w, err = strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(23)
	}
	rl := New(l, w)
	// Читаем остальные строки (пары ts key)
	for scanner.Scan() {
		line := scanner.Text()
		// Строка может быть пустой? По условию — нет, но на всякий случай.
		if line == "" {
			continue
		}

		parts = strings.Fields(line) // разделяем по пробелам
		if len(parts) < 2 {
			fmt.Printf("Skipping invalid line: %q\n", line)
			continue
		}
		ts, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Printf("Invalid ts: %v\n", err)
			continue
		}
		key := parts[1]
		rl.Limiting(key, int(ts))
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		os.Exit(100)
	}
}
