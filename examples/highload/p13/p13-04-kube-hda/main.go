package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	mn, mx, t, r0 int
}

func New(mn, mx, t, r0 int) Calculator {
	return Calculator{mn: mn, mx: mx, t: t, r0: r0}
}

func (c *Calculator) Calc(u int) int {
	diff := u - c.t
	if diff < 0 {
		diff *= -1
	}
	if diff*10 <= c.t {
		return c.r0
	}
	desired := (c.r0*u + c.t - 1) / c.t
	if desired < c.mn {
		desired = c.mn
	}

	if desired > c.mx {
		desired = c.mx
	}
	c.r0 = desired

	return desired
}

func main() {
	var (
		mn, mx, target, r0 int
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
	if len(parts) < 4 {
		fmt.Printf("Skipping invalid line: %q\n", 1)
		os.Exit(2)
	}
	mn, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(21)
	}

	mx, err = strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(22)
	}

	target, err = strconv.Atoi(parts[2])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(23)
	}

	r0, err = strconv.Atoi(parts[3])
	if err != nil {
		fmt.Printf("Invalid W: %v\n", err)
		os.Exit(24)
	}
	calc := New(mn, mx, target, r0)
	// Читаем остальные строки (пары ts key)
	for scanner.Scan() {
		var u int
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
		u, err = strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Invalid W: %v\n", err)
			os.Exit(41)
		}
		r := calc.Calc(u)
		fmt.Printf("%d\n", r)
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		os.Exit(100)
	}
}
