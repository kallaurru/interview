package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		b int64
		n int
	)

	_, err := fmt.Fscan(r, &b)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	// загружаем прыжки
	for i := 0; i < n; i++ {
		var name string
		var tms int64

		_, err = fmt.Fscan(r, &name, &tms)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}
		b -= tms
		if b >= 0 {
			fmt.Printf("%s OK %d\n", name, b)
		} else {
			fmt.Printf("%s DEADLINE_EXCEEDED %d\n", name, b+tms)
			break
		}
	}
	if b >= 0 {
		fmt.Printf("BUDGET_LEFT %d\n", b)
	}
}
