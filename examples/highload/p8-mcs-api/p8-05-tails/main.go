package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n   int
		err error
	)
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем первую строку — W
	if !scanner.Scan() {
		// Если ввода нет, завершаемся или обрабатываем ошибку
		fmt.Println("No input")
		return
	}
	firstLine := scanner.Text()
	n, err = strconv.Atoi(firstLine)
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		return
	}
	branch := make([]int, 0, (n+1)/2)
	funout := make([]int, 0, (n+1)/2)
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
