package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RLim struct {
	L, W     int
	scenario int // -1 - все отвергать, 0 - все пропускать
	state    int // 0 - хранилище отключено, 1 - хранилище доступно
	stor     map[string]RLItem
}

type RLItem struct {
	N int
	W int
}

func New(policy string, l, w int) *RLim {
	scenario := 0
	if policy == "closed" {
		scenario = -1
	}

	return &RLim{
		L:        l,
		W:        w,
		state:    1,
		scenario: scenario,
		stor:     make(map[string]RLItem, 256),
	}
}

func (rl *RLim) Limiting(t int, key string) {
	rl.print(rl.lim(t, key))
}

func (rl *RLim) Health(event string) {
	if event == "DOWN" {
		rl.state = 0
		return
	}

	rl.state = 1
	rl.stor = make(map[string]RLItem, 256)
}

func (rl *RLim) lim(t int, key string) bool {
	if rl.state == 0 {
		// хранилище не доступно
		return rl.limSpec()
	}
	// лимитируем по фиксированному окну
	if rl.L <= 0 {
		return false // нет лимита по запросам
	}

	data, ok := rl.stor[key]
	if !ok {
		data = RLItem{
			N: 1,
			W: t / rl.W,
		}
		rl.stor[key] = data
		return true
	}
	w := t / rl.W
	if w == data.W {
		if data.N < rl.L {
			data.N += 1
			rl.stor[key] = data

			return true
		}
		return false
	}
	data.N = 1
	data.W = w
	rl.stor[key] = data

	return true
}
func (rl *RLim) limSpec() bool {
	return rl.state+rl.scenario >= 0
}
func (rl *RLim) print(access bool) {
	if access {
		fmt.Printf("ALLOW\n")
		return
	}
	fmt.Printf("DENY\n")
}

func main() {
	var (
		l, w   int
		policy string
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
	if len(parts) < 3 {
		fmt.Printf("Skipping invalid line: %q\n", 1)
		os.Exit(2)
	}
	policy = parts[0]

	l, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Invalid L: %v\n", err)
		os.Exit(23)
	}
	w, err = strconv.Atoi(parts[2])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(24)
	}
	rl := New(policy, l, w)
	for scanner.Scan() {
		var (
			t          int
			event, key string
		)
		line := scanner.Text()
		// Строка может быть пустой? По условию — нет, но на всякий случай.
		if line == "" {
			continue
		}
		parts = strings.Fields(line) // разделяем по пробелам
		if len(parts) < 2 {
			fmt.Printf("Skipping invalid line: %q\n", line)
			continue
		}

		t, err = strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Invalid W: %v\n", err)
			os.Exit(23)
		}
		event = parts[1]
		if len(parts) == 2 {
			rl.Health(event)
			continue
		}
		key = parts[2]
		rl.Limiting(t, key)
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		os.Exit(100)
	}
}
