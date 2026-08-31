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
	var k, v, m int

	_, err := fmt.Fscan(r, &k, &v)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	pointsX := make([]Point, 0, k*v)
	pointsX1 := make([]Point, 0, (k+1)*v)
	names := make(map[string]struct{}, k)

	for i := 0; i < k; i++ {
		var name string

		_, err = fmt.Fscan(r, &name)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		_, ok := names[name]
		if ok {
			continue
		}
		pointsX = addNewNode(pointsX, name, v, i)   // i текущий номер ноды
		pointsX1 = addNewNode(pointsX1, name, v, i) // i текущий номер ноды
		names[name] = struct{}{}
	}

	var newNodeName string

	_, err = fmt.Fscan(r, &newNodeName)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(33)
	}

	_, ok := names[newNodeName]
	if !ok {
		pointsX1 = addNewNode(pointsX1, newNodeName, v, k) // k индекс новой ноды
	}
	_, err = fmt.Fscan(r, &m)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(33)
	}
	keys := make([]string, 0, m)

	for i := 0; i < m; i++ {
		var key string

		_, err = fmt.Fscan(r, &key)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(4)
		}

		keys = append(keys, key)
	}

	changed := 0
	for i := 0; i < len(keys); i++ {
		pointX := findNode(pointsX, keys[i])
		pointX1 := findNode(pointsX1, keys[i])
		if pointX.Node != pointX1.Node {
			changed++
		}
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

func addNewNode(points []Point, name string, v, k int) []Point {
	for i := 0; i < v; i++ {
		label := fmt.Sprintf("%s#%d", name, i)
		p := Point{
			Node:  k,
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
