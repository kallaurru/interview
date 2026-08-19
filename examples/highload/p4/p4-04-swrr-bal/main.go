package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, n int

	_, err := fmt.Fscan(r, &k)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	bW := make(map[string]int, k)
	bC := make(map[string]int, k)
	bs := make([]string, k)
	allW := 0

	for i := 0; i < k; i++ {
		var name string
		var w int

		_, err = fmt.Fscan(r, &name, &w)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}
		bW[name] = w
		allW += w
		bs[i] = name
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(33)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", balance(bW, bC, bs, allW))
	}
}

func balance(bW, bc map[string]int, bs []string, allW int) string {
	selectedName := bs[0]
	selectedCounter := bc[selectedName] + bW[selectedName]
	bc[selectedName] = selectedCounter

	for i := 1; i < len(bs); i++ {
		name := bs[i]
		tmp := bc[name] + bW[name]
		if tmp > selectedCounter {
			selectedName = name
			selectedCounter = tmp
		}
		// обновить счетчик
		bc[name] = tmp
	}
	bc[selectedName] = selectedCounter - allW

	return selectedName
}
