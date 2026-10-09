package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	pos     uint32
	shardID int
	nodeID  int
}

func New(shardID, nodeID int) Point {
	return Point{
		shardID: shardID,
		nodeID:  nodeID,
	}
}

func (p Point) Calc(hf func(key []byte) uint32) Point {
	return Point{
		shardID: p.shardID,
		nodeID:  p.nodeID,
		pos:     hf([]byte(p.Label())),
	}
}

func (p Point) Label() string {
	return fmt.Sprintf("shard%d#%d", p.shardID, p.nodeID)
}

func (p Point) Pos() uint32 {
	return p.pos
}

func (p Point) ShardID() string {
	return fmt.Sprintf("shard%d", p.shardID)
}

func main() {
	var (
		n, v int
	)
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем первую строку
	if !scanner.Scan() {
		// Если ввода нет, завершаемся или обрабатываем ошибку
		fmt.Println("No input")
		os.Exit(1)
	}
	firstLine := scanner.Text()
	parts := strings.Fields(firstLine) // разделяем по пробелам
	if len(parts) < 2 {
		fmt.Printf("Skipping invalid line: %q\n", 1)
		os.Exit(2)
	}
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(22)
	}

	v, err = strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(23)
	}
	points := makePointsArea(n, v)
	// Читаем остальные строки (пары ts key)
	for scanner.Scan() {
		line := scanner.Text()
		// Строка может быть пустой? По условию — нет, но на всякий случай.
		if line == "" {
			continue
		}
		parts = strings.Fields(line) // разделяем по пробелам
		if len(parts) < 1 {
			fmt.Printf("Skipping invalid line: %q\n", line)
			continue
		}
		p := findNode(points, parts[0])
		fmt.Printf("%s %s\n", parts[0], p.ShardID())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		os.Exit(100)
	}
}

func fnv(key []byte) uint32 {
	const (
		basis uint32 = 2166136261
		h1    uint32 = 16777619
	)

	var out uint32
	out = basis
	for _, b := range key {
		out ^= uint32(b)
		out = (out * h1) % (1<<32 - 1)
	}

	return out
}

func makePointsArea(n, v int) []Point {
	out := make([]Point, 0, n*v)
	for i := 0; i < n; i++ {
		for j := 0; j < v; j++ {
			p := New(i, j)
			p = p.Calc(fnv)
			out = append(out, p)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].Pos() != out[j].Pos() {
			return out[i].Pos() < out[j].Pos()
		}
		return out[i].Label() < out[j].Label()
	})

	return out
}

func findNode(points []Point, key string) Point {
	kp := fnv([]byte(key))
	idx := sort.Search(len(points), func(i int) bool { return points[i].Pos() >= kp })
	if idx == len(points) {
		idx = 0 // прошли конец кольца — заворачиваемся
	}

	return points[idx]
}
