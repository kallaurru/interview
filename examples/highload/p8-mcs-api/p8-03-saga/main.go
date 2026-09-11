package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const opFAIL = "FAIL"

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	saga := make([]string, 0, n-1)
	compS := make([]string, 0, n-1)
	saga = append(saga, "")
	compS = append(compS, "")

	for i := 0; i < n; i++ {
		var step, comp string

		_, err = fmt.Fscan(r, &step, &comp)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(1)
		}
		saga = append(saga, step)
		compS = append(compS, comp)
	}
	var (
		lastStep string
		step     int
	)
	line, err := r.ReadString('\n')
	line, err = r.ReadString('\n')

	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(11)
	}
	fields := strings.Fields(line)
	if len(fields) > 0 {
		lastStep = fields[0]
	}
	if len(fields) > 1 {
		step, err = strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(22)
		}
	}
	for i := 1; i < len(saga); i++ {
		if lastStep == opFAIL && step == i {
			fmt.Printf("FAIL %s\n", saga[i])
			break
		}
		fmt.Printf("EXEC %s\n", saga[i])
	}
	if lastStep == opFAIL {
		for i := step - 1; i >= 1; i-- {
			fmt.Printf("COMP %s\n", compS[i])
		}
	}

	if lastStep == opFAIL {
		fmt.Printf("%s\n", "RESULT COMPENSATED")
	} else {
		fmt.Printf("%s\n", "RESULT COMMITTED")
	}
}
