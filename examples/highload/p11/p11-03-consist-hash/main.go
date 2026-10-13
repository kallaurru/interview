package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func fnv1a64(s string) uint64 {
	h := uint64(14695981039346269339)
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211
	}
	return h
}

type vnode struct {
	hash uint64
	node string
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
	V, _ := strconv.Atoi(parts[1])

	nodes := make([]string, N)
	for i := 0; i < N; i++ {
		if !sc.Scan() {
			return
		}
		nodes[i] = sc.Text()
	}

	// Build the ring of virtual nodes
	vnodes := make([]vnode, 0, N*V)
	for _, name := range nodes {
		for i := 0; i < V; i++ {
			s := name + "#" + strconv.Itoa(i)
			vnodes = append(vnodes, vnode{hash: fnv1a64(s), node: name})
		}
	}

	sort.Slice(vnodes, func(i, j int) bool {
		if vnodes[i].hash != vnodes[j].hash {
			return vnodes[i].hash < vnodes[j].hash
		}
		return vnodes[i].node < vnodes[j].node
	})

	counts := make(map[string]int, N)

	for sc.Scan() {
		key := sc.Text()
		h := fnv1a64(key)

		// First vnode with hash >= h
		idx := sort.Search(len(vnodes), func(i int) bool {
			return vnodes[i].hash >= h
		})
		// Wrap around the ring
		if idx == len(vnodes) {
			idx = 0
		}
		counts[vnodes[idx].node]++
	}

	var sb strings.Builder
	for _, name := range nodes {
		fmt.Fprintf(&sb, "%s=%d\n", name, counts[name])
	}
	fmt.Print(sb.String())
}
