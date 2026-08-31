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
		n              int64
		key            string
		allMin, allMax int64
	)
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем первую строку — n
	if !scanner.Scan() {
		// Если ввода нет, завершаемся или обрабатываем ошибку
		fmt.Println("No input")
		return
	}
	firstLine := scanner.Text()
	n, err := strconv.ParseInt(firstLine, 10, 64)
	if err != nil {
		fmt.Printf("Invalid n: %v\n", err)
		return
	}
	stat := make(map[uint32]int64, n)
	// Читаем остальные строки (пары ts key)
	for scanner.Scan() {
		line := scanner.Text()
		// Строка может быть пустой? По условию — нет, но на всякий случай.
		if line == "" {
			continue
		}
		key = strings.Trim(line, " ")
		part := hasher(key) % uint32(n)

		stat[part] += 1
	}
	allMin = -1
	for i := 0; i < int(n); i++ {
		counter := stat[uint32(i)]
		fmt.Printf("partition %d: %d\n", i, counter)
		if allMin == -1 {
			allMin = counter
		}
		if counter > allMax {
			allMax = counter
		}
		if counter < allMin {
			allMin = counter
		}
	}

	fmt.Printf("skew=%d\n", allMax-allMin)
	// Проверяем ошибки сканера (если были)
	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}
}

func hasher(key string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		h ^= uint32(key[i])
		h *= 16777619
	}
	return h
}
