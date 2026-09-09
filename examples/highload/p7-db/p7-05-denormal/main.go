package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	cmdFollow = "FOLLOW"
	cmdUnf    = "UNFOLLOW"
	cmdPost   = "POST"
)

type Manager struct {
	followers map[string]int
	total     int
	max       int
}

func New(n int) *Manager {
	return &Manager{
		followers: make(map[string]int, (n+1)/2),
		total:     0,
		max:       0,
	}
}

func (m *Manager) Op(op string, opts ...string) error {
	switch op {
	case cmdFollow:
		if len(opts) < 2 {
			return errors.New("count of opts must be 2 or more")
		}
		m.followers[opts[1]] += 1
	case cmdUnf:
		if len(opts) < 2 {
			return errors.New("count of opts must be 2 or more")
		}
		m.followers[opts[1]] -= 1
	case cmdPost:
		if len(opts) < 1 {
			return errors.New("count of opts must be 1 or more")
		}
		followers := m.followers[opts[0]]
		m.total += followers
		if m.max < followers {
			m.max = followers
		}
	}

	return nil
}

func (m *Manager) Total() int {
	return m.total
}

func (m *Manager) Max() int {
	return m.max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем первую строку — n
	if !scanner.Scan() {
		// Если ввода нет, завершаемся или обрабатываем ошибку
		fmt.Println("No input")
		return
	}
	firstLine := scanner.Text()
	n, err := strconv.Atoi(firstLine)
	if err != nil {
		fmt.Printf("Invalid r: %v\n", err)
		return
	}
	mngr := New(n)
	for scanner.Scan() {

		line := scanner.Text()
		// Строка может быть пустой? По условию — нет, но на всякий случай.
		if line == "" {
			continue
		}
		items := strings.Fields(line)
		if len(items) < 2 {
			fmt.Println("Bad line. Has less 2 param")
			return
		}
		err = mngr.Op(items[0], items[1:]...)
		if err != nil {
			fmt.Printf("Mngr err: %v\n", err)
			return
		}
	}

	fmt.Printf("%d %d\n", mngr.Total(), mngr.Max())
}
