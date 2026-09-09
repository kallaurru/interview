package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	Pos   uint32
	Node  int
	Label string
	Name  string
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, v, nOld, nNew int

	_, err := fmt.Fscan(r, &nOld, &nNew, &v, &k)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	pointsX := makeRing(nNew, v)
	pointsXOld := makeRing(nOld, v)

	movedMod := 0
	movedRing := 0

	for i := 0; i < k; i++ {
		var key string

		_, err = fmt.Fscan(r, &key)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}

		// schema mod
		movedMod += calcSchemaMod(key, nOld, nNew)
		movedRing += calcSchemaRing(key, pointsXOld, pointsX)
	}

	fmt.Printf("%d %d\n", movedMod, movedRing)
}

func fnv32a(s string) uint32 {
	h := uint32(2166136261)       // offset basis
	for i := 0; i < len(s); i++ { // перебор БАЙТОВ строки
		h ^= uint32(s[i]) // сначала XOR…
		h *= 16777619     // …потом умножение (uint32 сам даёт mod 2^32)
	}
	return h
}

func calcSchemaMod(key string, nOld, nNew int) int {
	shOld := fnv32a(key) % uint32(nOld)
	shNew := fnv32a(key) % uint32(nNew)
	if shOld == shNew {
		return 0
	}
	return 1
}

func calcSchemaRing(key string, old, new []Point) int {
	oldP := findNode(old, key)
	oldN := findNode(new, key)

	if oldN.Node == oldP.Node {
		return 0
	}
	return 1
}

func findNode(points []Point, key string) Point {
	kp := fnv32a(key)
	idx := sort.Search(len(points), func(i int) bool { return points[i].Pos >= kp })
	if idx == len(points) {
		idx = 0 // прошли конец кольца — заворачиваемся
	}

	return points[idx]
}

func makeRing(shards int, v int) []Point {
	out := make([]Point, 0, shards*v)
	for s := 0; s < shards; s++ {
		for vk := 0; vk < v; vk++ {
			label := fmt.Sprintf("shard-%d#%d", s, vk)
			p := Point{
				Node:  s,
				Name:  "",
				Label: label,
				Pos:   fnv32a(label),
			}
			out = append(out, p)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Pos != out[j].Pos {
			return out[i].Pos < out[j].Pos
		}
		return out[i].Label < out[j].Label // tie-break по метке
	})

	return out
}
