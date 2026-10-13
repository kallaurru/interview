package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	var K int
	var T int64

	if !sc.Scan() {
		return
	}
	fmt.Sscan(sc.Text(), &K, &T)

	var pushAll int64
	var pushHybrid int64

	for i := 0; i < K; i++ {
		if !sc.Scan() {
			break
		}
		var followers, posts int64
		fmt.Sscan(sc.Text(), &followers, &posts)

		product := followers * posts
		pushAll += product

		if followers <= T {
			pushHybrid += product
		}
	}

	var saved int64
	if pushAll != 0 {
		saved = (pushAll - pushHybrid) * 100 / pushAll
	}

	fmt.Printf("push_all=%d\n", pushAll)
	fmt.Printf("push_hybrid=%d\n", pushHybrid)
	fmt.Printf("saved=%d%%\n", saved)
}
