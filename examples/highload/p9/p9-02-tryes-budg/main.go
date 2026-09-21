package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var b, base, gap, attempts, leftover, elapsed int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &b, &base, &gap)
	if err != nil {
		os.Exit(1)
	}
	k := 0
	elapsed += base
	if b-base < 0 {
		fmt.Printf("%d %d\n", attempts, b)
		return
	}
	attempts++
	for {
		k++
		pause := base * (1 << k)
		if b-elapsed-gap-pause >= 0 {
			attempts++
			elapsed += gap + pause
			continue
		}
		break
	}
	leftover = b - elapsed
	fmt.Printf("%d %d\n", attempts, leftover)
}
