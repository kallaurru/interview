package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deduplicate struct {
	keys map[string]int64
	w    int64
}

func New(w int64) *Deduplicate {
	return &Deduplicate{
		keys: make(map[string]int64, 128),
		w:    w,
	}
}

func (dd *Deduplicate) Once(key string, ts int64) bool {
	lastTs, seen := dd.keys[key]
	if !seen || ts-lastTs >= dd.w {
		dd.keys[key] = ts // приняли
		return true
	}
	return false
}

func main() {
	var (
		w, accepted, duplicated, ts int64
		key                         string
	)
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем первую строку — W
	if !scanner.Scan() {
		// Если ввода нет, завершаемся или обрабатываем ошибку
		fmt.Println("No input")
		return
	}
	firstLine := scanner.Text()
	w, err := strconv.ParseInt(firstLine, 10, 64)
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		return
	}

	dd := New(w)
	// Читаем остальные строки (пары ts key)
	for scanner.Scan() {
		line := scanner.Text()
		// Строка может быть пустой? По условию — нет, но на всякий случай.
		if line == "" {
			continue
		}
		parts := strings.Fields(line) // разделяем по пробелам
		if len(parts) < 2 {
			fmt.Printf("Skipping invalid line: %q\n", line)
			continue
		}
		ts, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Printf("Invalid ts: %v\n", err)
			continue
		}
		key = parts[1]

		if dd.Once(key, ts) {
			accepted++
			fmt.Printf("ACCEPT %s\n", key)
		} else {
			duplicated++
			fmt.Printf("DUPLICATE %s\n", key)
		}
	}

	fmt.Printf("accepted=%d duplicates=%d\n", accepted, duplicated)
	// Проверяем ошибки сканера (если были)
	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}
}
