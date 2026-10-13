package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !sc.Scan() {
		return
	}
	P, err := strconv.Atoi(strings.TrimSpace(sc.Text()))
	if err != nil {
		return
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	total := 0
	sampled := 0

	for sc.Scan() {
		traceID := strings.TrimSpace(sc.Text())
		if traceID == "" {
			continue
		}
		total++

		val, err := strconv.ParseUint(traceID, 16, 64)
		if err != nil {
			continue
		}

		if val%100 < uint64(P) {
			fmt.Fprintln(out, traceID)
			sampled++
		}
	}

	fmt.Fprintf(out, "sampled %d/%d\n", sampled, total)
}
