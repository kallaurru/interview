package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var deadline, n int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &deadline)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	for i := 0; i < n; i++ {
		var (
			net, work int
			name      string
		)
		_, err = fmt.Fscan(r, &name, &net, &work)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}
		deadline -= net
		if deadline <= 0 {
			fmt.Printf("TIMEOUT %s network\n", name)
			break
		}
		deadline -= work
		if deadline <= 0 {
			fmt.Printf("TIMEOUT %s work\n", name)
			break
		}
	}

	if deadline > 0 {
		fmt.Printf("%s %d\n", "OK", deadline)
	}

}
