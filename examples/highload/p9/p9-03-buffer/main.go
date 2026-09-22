package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var c, s, n int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &c, &s)
	if err != nil {
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(2)
	}

	var dropped, maxq, buf int

	for i := 0; i < n; i++ {
		var takt int

		_, err = fmt.Fscan(r, &takt)
		if err != nil {
			os.Exit(33)
		}
		buf += takt
		if buf > c {
			dropped += buf - c
			buf -= buf - c
		}
		maxq = max(maxq, buf)
		leave := min(buf, s)
		buf -= leave
	}

	fmt.Printf("%d %d %d\n", dropped, maxq, buf)
}
