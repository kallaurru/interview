package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Container struct {
	Name       string
	LastReason string
	LastExit   string
	Restarts   int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	var containers []*Container
	var current *Container
	section := ""

	for scanner.Scan() {
		raw := scanner.Text()
		trimmedRight := strings.TrimRight(raw, " \t\r\n")
		indent := 0
		for indent < len(trimmedRight) && trimmedRight[indent] == ' ' {
			indent++
		}
		content := trimmedRight[indent:]

		// Строка с именем контейнера: отступ ровно 2 пробела,
		// без пробелов внутри, заканчивается двоеточием.
		if indent == 2 && strings.HasSuffix(content, ":") {
			name := strings.TrimSuffix(content, ":")
			if name != "" && !strings.Contains(name, " ") {
				current = &Container{Name: name}
				containers = append(containers, current)
				section = ""
				continue
			}
		}

		if current == nil {
			continue
		}

		// Отслеживаем, в каком блоке находимся.
		if strings.HasPrefix(content, "Last State:") {
			section = "laststate"
			continue
		}
		if strings.HasPrefix(content, "State:") {
			section = "state"
			continue
		}

		// Reason: относится к последнему State / Last State.
		if strings.HasPrefix(content, "Reason:") {
			val := strings.TrimSpace(strings.TrimPrefix(content, "Reason:"))
			if section == "laststate" {
				current.LastReason = val
			}
			continue
		}

		// Exit Code: аналогично.
		if strings.HasPrefix(content, "Exit Code:") {
			val := strings.TrimSpace(strings.TrimPrefix(content, "Exit Code:"))
			if section == "laststate" {
				current.LastExit = val
			}
			continue
		}

		// Restart Count: количество рестартов контейнера.
		if strings.HasPrefix(content, "Restart Count:") {
			val := strings.TrimSpace(strings.TrimPrefix(content, "Restart Count:"))
			n, err := strconv.Atoi(val)
			if err == nil {
				current.Restarts = n
			}
			continue
		}
	}

	total := 0
	maxRestarts := -1
	topName := ""
	anyOutput := false

	for _, c := range containers {
		total += c.Restarts
		if c.Restarts > maxRestarts {
			maxRestarts = c.Restarts
			topName = c.Name
		}
		if c.Restarts > 0 {
			reason := c.LastReason
			if reason == "" {
				reason = "-"
			}
			exit := c.LastExit
			if exit == "" {
				exit = "-"
			}
			fmt.Printf("%s last=%s exit=%s restarts=%d\n", c.Name, reason, exit, c.Restarts)
			anyOutput = true
		}
	}

	if !anyOutput {
		fmt.Println("no restarts")
		return
	}

	fmt.Printf("total=%d top=%s\n", total, topName)
}
