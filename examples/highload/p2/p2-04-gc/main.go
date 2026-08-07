package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, gc, live, temp, target int
	targetF := func(live, gc int) int {
		return live * (100 + gc) / 100
	}
	_, err := fmt.Fscan(r, &gc, &live)
	if err != nil {
		os.Exit(1)
	}
	target = targetF(live, gc)

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(2)
	}
	if n <= 0 {
		fmt.Printf("%d %d %d\n", 0, live, targetF(live, gc))
		os.Exit(0)
	}
	countGC := 0
	for i := 0; i < n; i++ {
		var flag string
		var size int

		_, err = fmt.Fscan(r, &size, &flag)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		need := dispatchE(size, flag, &live, &temp, target)
		if need {
			temp = 0
			target = targetF(live, gc)
			countGC++
		}
	}

	fmt.Printf("%d %d %d\n", countGC, live, target)
}

func dispatchE(size int, flg string, live, temp *int, target int) bool {
	if flg == "L" {
		*live += size
	}
	if flg == "T" {
		*temp += size
	}
	if *live+*temp >= target {
		// требуется сборка
		return true
	}

	return false
}
