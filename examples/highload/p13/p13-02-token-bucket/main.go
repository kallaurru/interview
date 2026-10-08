package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TBucket struct {
	C, R int
	stor map[string]TBItem
}

type TBItem struct {
	T int
	L int64
}

func New(c, r int) *TBucket {
	return &TBucket{C: c, R: r, stor: make(map[string]TBItem, 256)}
}

func (tb *TBucket) Limiting(key string, ts int64) {
	val, ok := tb.lim(key, ts)

	tb.print(val, ok)
}

func (tb *TBucket) lim(key string, ts int64) (int, bool) {
	if tb.C == 0 {
		return 0, false
	}
	item, ok := tb.stor[key]

	if !ok {
		item = TBItem{T: tb.C - 1, L: ts}
		tb.stor[key] = item
		return item.T, true
	}
	add := int(int64(item.T) + (ts-item.L)*int64(tb.R))
	tokens := minLoc(tb.C, add)
	if tokens < 1 {
		return 0, false
	}
	item.T = tokens - 1
	item.L = ts
	tb.stor[key] = item

	return item.T, true
}

func (tb *TBucket) print(count int, access bool) {
	if access {
		fmt.Printf("ALLOW %d\n", count)
		return
	}
	fmt.Printf("DENY %d\n", 0)
}

func main() {
	var (
		c, r int
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
	c, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(22)
	}

	r, err = strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(23)
	}

	tb := New(c, r)
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

		tb.Limiting(key, ts)
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		os.Exit(100)
	}
}

func minLoc(a, b int) int {
	if a < b {
		return a
	}
	return b
}
