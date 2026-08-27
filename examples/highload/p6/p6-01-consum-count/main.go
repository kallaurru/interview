package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		l, m, b, d int64
	)
	_, err := fmt.Fscan(r, &l, &m, &b, &d)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	/*	_, err = fmt.Fscan(r, &n)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
	*/
	n := int64(0)

	if l == 0 && b == 0 {
		n = 1
		fmt.Printf("consumers=%d\n", n)
		return
	}
	if b == 0 {
		rem := l % m
		if rem > 0 {
			n = l/m + 1
		} else {
			n = l / m
		}
		fmt.Printf("consumers=%d\n", n)
		return
	}
	// 90 40 1000 10 l m b d
	// b > 0
	// n >= (l/m) + (b/(d*m)

	n = l / m // для разбора потока
	if n*m <= l {
		n++
	}
	for d*(n*m-l) < b {
		n++
	}

	fmt.Printf("consumers=%d\n", n)
}
