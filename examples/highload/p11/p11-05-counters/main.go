package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fnv1a32(s string) uint32 {
	h := uint32(2166136261)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !sc.Scan() {
		return
	}
	parts := strings.Fields(sc.Text())
	if len(parts) < 2 {
		return
	}
	N, _ := strconv.Atoi(parts[0])
	B := 0
	fmt.Sscan(parts[1], &B)

	counts := make([]int64, N)
	buffers := make([]int64, N)
	flushes := make([]int64, N)
	n32 := uint32(N)
	var M int64

	for sc.Scan() {
		eventID := strings.TrimSpace(sc.Text())
		if eventID == "" {
			continue
		}
		M++
		slot := fnv1a32(eventID) % n32
		counts[slot]++
		buffers[slot]++
		if buffers[slot] >= int64(B) {
			flushes[slot]++
			buffers[slot] = 0
		}
	}

	// Flush remaining non-empty buffers
	var F int64
	for i := 0; i < N; i++ {
		if buffers[i] > 0 {
			flushes[i]++
			buffers[i] = 0
		}
		F += flushes[i]
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < N; i++ {
		fmt.Fprintf(out, "slot %d: count=%d flushes=%d\n", i, counts[i], flushes[i])
	}

	var saved int64
	if M > 0 {
		saved = (M - F) * 100 / M
	}
	fmt.Fprintf(out, "naive=%d batched=%d saved=%d%%\n", M, F, saved)
}
