package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var e, m int
	count := int64(0)
	_, err := fmt.Fscan(r, &e, &m)
	if err != nil {
		os.Exit(1)
	}
	if m <= 0 {
		os.Exit(11)
	}
	for i := 0; i < m; i++ {
		var num int
		_, err = fmt.Fscan(r, &num)
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		count += int64(num)
		if count > int64(e) {
			fmt.Printf("%d\n", i+1)
			os.Exit(0)
		}
	}
	fmt.Printf("%s\n", "WITHIN_BUDGET")
}
