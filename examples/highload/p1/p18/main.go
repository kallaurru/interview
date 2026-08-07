package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var mn, mx, target, r0, steps int

	_, err := fmt.Fscan(r, &mn, &mx, &target, &r0)
	if err != nil {
		os.Exit(1)
	}
	_, err = fmt.Fscan(r, &steps)
	if err != nil {
		os.Exit(2)
	}
	for i := 0; i < steps; i++ {
		var val int
		_, err = fmt.Fscan(r, &val)

		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		desired := (r0*val + target - 1) / target
		r0 = calcR(desired, mn, mx)
		fmt.Printf("%d\n", r0)
	}
}

func calcR(r, mn, mx int) int {
	if r < mn {
		return mn
	}
	if r > mx {
		return mx
	}

	return r
}
