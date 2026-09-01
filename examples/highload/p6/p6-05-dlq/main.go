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
		r, dlqTotal  int
		id, outcomes string
	)
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем первую строку — n
	if !scanner.Scan() {
		// Если ввода нет, завершаемся или обрабатываем ошибку
		fmt.Println("No input")
		return
	}
	firstLine := scanner.Text()
	r, err := strconv.Atoi(firstLine)
	if err != nil {
		fmt.Printf("Invalid r: %v\n", err)
		return
	}

	// Читаем остальные строки (пары ts key)
	for scanner.Scan() {
		line := scanner.Text()
		// Строка может быть пустой? По условию — нет, но на всякий случай.
		if line == "" {
			continue
		}
		items := strings.Fields(line)
		if len(items) < 2 {
			continue
		}
		id = items[0]
		outcomes = items[1]
		countTry := r + 1
		ok := false
		for i := 0; i < len(outcomes); i++ {
			if countTry <= 0 {
				break
			}
			if outcomes[i] == 'S' {
				ok = true
				fmt.Printf("%s ok attempts=%d\n", id, i+1)
				break
			}
			countTry--
		}
		if !ok {
			dlqTotal++
			fmt.Printf("%s dlq\n", id)
		}
	}

	fmt.Printf("dlq_total=%d\n", dlqTotal)
}
