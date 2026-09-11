package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	var sumV, sumS, n int

	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	for i := 0; i < n; i++ {
		var item uint64
		_, err := fmt.Fscan(r, &item)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}

		val := calculator(item)
		fmt.Printf("%d\n", val)
		valS := len(fmt.Sprintf("%d", item))
		sumV += val
		sumS += valS
	}

	fmt.Printf("%d %d\n", sumV, sumS)
}

// Возвращаем количество байт по varint кодировке
func calculator(val uint64) int {
	if val == 0 {
		return 1
	}
	minBits := bits.Len64(val)
	count := minBits / 7
	rem := minBits % 7
	if rem > 0 {
		return count + 1
	}
	return count
}
