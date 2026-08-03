package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, m int

	_, err := fmt.Fscan(r, &k, &m)
	if err != nil {
		os.Exit(1)
	}

	names := make(map[int]string)
	oldValues := make([]int, m)
	newValues := make([]int, m)

	//загрузили предыдущие значения
	for i := 0; i < k; i++ {
		idx := i % k
		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		fields := strings.Fields(string(line))
		if len(fields) < m+1 {
			fmt.Printf("Линия должна иметь #%d: полей %v\n", m+1, err)
			os.Exit(3)
		}
		names[idx] = fields[0]
		for mm := 1; mm <= m; mm++ {
			val, err := strconv.Atoi(fields[mm])
			if err != nil {
				fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
				os.Exit(3)
			}
			oldValues[mm-1] = val
		}
	}
	sort.Ints(oldValues)

	// загружаем новые
	for i := k; i < k*2; i++ {
		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		fields := strings.Fields(string(line))
		if len(fields) < m+1 {
			fmt.Printf("Линия должна иметь #%d: полей %v\n", m+1, err)
			os.Exit(3)
		}
		for mm := 1; mm <= m; mm++ {
			val, err := strconv.Atoi(fields[mm])
			if err != nil {
				fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
				os.Exit(3)
			}
			newValues[mm-1] = val
		}
	}
	sort.Ints(newValues)
}
