package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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
	fanout := make([]int, 0, (n+1)/2)
	// Читаем остальные строки (пары ts key)
	for scanner.Scan() {
		var fnOutMax int
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
		k, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Printf("Invalid ts: %v\n", err)
			continue
		}
		if len(parts[1:]) < int(k) {
			fmt.Printf(": %v\n", err)
			os.Exit(333)
		}

		for i := 1; i <= int(k); i++ {
			val, err := strconv.Atoi(parts[i])
			if err != nil {
				fmt.Printf("Invalid ts: %v\n", err)
				continue
			}
			branch = append(branch, val)
			if fnOutMax < val {
				fnOutMax = val
			}
		}

		fanout = append(fanout, fnOutMax)
	}
	slices.Sort(branch)
	slices.Sort(fanout)
	pb50Idx := math.Ceil(float64(50*len(branch)) / float64(100))
	pb99Idx := math.Ceil(float64(99*len(branch)) / float64(100))

	pf50Idx := math.Ceil(float64(50*len(fanout)) / float64(100))
	pf99Idx := math.Ceil(float64(99*len(fanout)) / float64(100))

	pb50 := branch[int(pb50Idx)-1]
	pb99 := branch[int(pb99Idx)-1]
	pf50 := fanout[int(pf50Idx)-1]
	pf99 := fanout[int(pf99Idx)-1]

	fmt.Printf("BRANCH %d %d\n", pb50, pb99)
	fmt.Printf("FANOUT %d %d\n", pf50, pf99)
}
