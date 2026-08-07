package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 4*1024*1024) // строки до 2 МБ
	sc.Scan()                                       // первая строка: параметры
	parts := strings.Fields(sc.Text())
	if len(parts) != 2 {
		os.Exit(1)
	}
	sc.Scan() // вторая строка: до 10^6 символов
	bits := strings.TrimSpace(sc.Text())
	failed := strings.Count(bits, "0")
	s, err := strconv.Atoi(parts[0])
	if err != nil {
		os.Exit(2)
	}
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		os.Exit(21)
	}

	a, l, ok := errorBudgetCalc(s, n, failed)
	fmt.Printf("allowed=%d\n", a)
	fmt.Printf("failed=%d\n", failed)
	fmt.Printf("left=%d\n", l)
	if ok {
		fmt.Printf("OK")
	} else {
		fmt.Printf("VIOLATED")
	}
}

func errorBudgetCalc(s, n, f int) (int, int, bool) {
	allowed := n * (100000 - s) / 100000
	left := allowed - f
	return allowed, left, left >= 0
}
