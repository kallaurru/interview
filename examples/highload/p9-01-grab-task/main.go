package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var w, n int

	_, err := fmt.Fscan(r, &w, &n)
	if err != nil {
		os.Exit(1)
	}

	numbers := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var val int
		_, err = fmt.Fscan(r, &val)
		if err != nil {
			os.Exit(2)
		}
		numbers = append(numbers, val)
	}

	makespan := 0
	idle := 0

	fmt.Printf("%d %d\n", makespan, idle)
}
