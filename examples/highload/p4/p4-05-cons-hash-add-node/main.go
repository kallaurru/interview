package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	Pos   uint32
	Idx   int
	Label string
	Name  string
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, v, m int

	_, err := fmt.Fscan(r, &k, &v)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	points := make([]Point, 0, (k+1)*v)

	for i := 0; i < k; i++ {
		var name string

		_, err = fmt.Fscan(r, &name)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		points = addNewNode(points, name, v)
	}

	var newNodeName string

	_, err = fmt.Fscan(r, &newNodeName)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(33)
	}

	_, err = fmt.Fscan(r, &m)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(33)
	}
	oldVersion := make(map[string]Point, m)
	keys := make([]string, 0, m)

	for i := 0; i < m; i++ {
		var key string

		_, err = fmt.Fscan(r, &key)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(4)
		}

		keys = append(keys, key)
		point := findNode(points, key)
		oldVersion[key] = point
	}

	points = addNewNode(points, newNodeName, v)
	newVersion := make(map[string]Point, m)

	for i := 0; i < len(keys); i++ {
		point := findNode(points, keys[i])
		newVersion[keys[i]] = point
	}

	changed := 0
	for key, node := range newVersion {
		oldNode, ok := oldVersion[key]
		if !ok {
			changed++
			continue
		}

		if oldNode.Name == node.Name {
			continue
		}
		changed++
	}
	fmt.Printf("%d\n", changed)
}

func fnv32a(s string) uint32 {
	h := uint32(2166136261)       // offset basis
	for i := 0; i < len(s); i++ { // перебор БАЙТОВ строки
		h ^= uint32(s[i]) // сначала XOR…
		h *= 16777619     // …потом умножение (uint32 сам даёт mod 2^32)
	}
	return h
}

func findNode(points []Point, key string) Point {
	kp := fnv32a(key)
	idx := sort.Search(len(points), func(i int) bool { return points[i].Pos >= kp })
	if idx == len(points) {
		idx = 0 // прошли конец кольца — заворачиваемся
	}

	return points[idx]
}

func addNewNode(points []Point, name string, v int) []Point {
	for ii := 0; ii < v; ii++ {
		label := fmt.Sprintf("%s#%d", name, ii)
		p := Point{
			Idx:   ii,
			Name:  name,
			Label: label,
			Pos:   fnv32a(label),
		}
		points = append(points, p)
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].Pos != points[j].Pos {
			return points[i].Pos < points[j].Pos
		}
		return points[i].Label < points[j].Label // tie-break по метке
	})

	return points
}
