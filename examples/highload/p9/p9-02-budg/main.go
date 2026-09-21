package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var b, n int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &b)
	if err != nil {
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(2)
	}
	elapsed := 0
	for i := 0; i < n; i++ {
		var (
			stage string
			bb    int
		)

		_, err = fmt.Fscan(r, &stage, &bb)
		if err != nil {
			os.Exit(3)
		}
		if b-elapsed <= 0 {
			fmt.Printf("canceled %s\n", stage)
			return
		}

		if b-elapsed < bb {
			fmt.Printf("timeout %s\n", stage)
			return
		}

		elapsed += bb
		fmt.Printf("ok %s %d\n", stage, b-elapsed)
	}
	fmt.Printf("done %d\n", b-elapsed)
}
