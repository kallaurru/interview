package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type RLim struct {
	L, W int
	stor map[string]RLItem
}

type RLItem struct {
	N int
	S int64
}

func New(l, w int) *RLim {
	return &RLim{L: l, W: w, stor: make(map[string]RLItem, 256)}
}

func (rl *RLim) Limiting(key string) {
	val, ok := rl.lim(key)

	rl.print(val, ok)
}

func (rl *RLim) lim(key string) (int, bool) {
	item, ok := rl.stor[key]
	moment := time.Now().Unix()
	if !ok {
		item = RLItem{N: 1, S: moment}
		rl.stor[key] = item
		return rl.L - item.N, true
	}
	// найдены данные
	if item.N < rl.L {
		item.N++
		rl.stor[key] = item
		return rl.L - item.N, true
	}
	endW := item.S + int64(rl.W)
	if endW > moment {
		return 0, false
	}
	// открытие нового окна
	item.S = moment
	item.N = 1

	return rl.L - item.N, true
}

func (rl *RLim) print(count int, access bool) {
	if access {
		fmt.Printf("200 remaining=%d\n", count)
		return
	}
	fmt.Printf("429 remaining=0 retry_after=%d\n", rl.W)
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
		if len(parts) < 1 {
			fmt.Printf("Skipping invalid line: %q\n", line)
			continue
		}
		rl.Limiting(parts[0])
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		os.Exit(100)
	}
}
