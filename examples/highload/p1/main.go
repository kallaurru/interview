package main

import (
	"bufio"
	"fmt"
	"github.com/kallaurru/interview/polygon/highload/p1"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	val, _, err := in.ReadLine()
	if err != nil {
		os.Exit(1)
	}
	str := string(val)
	fields := strings.Split(str, " ")
	if len(fields) != 4 {
		os.Exit(2)
	}
	dau, err := strconv.Atoi(fields[0])
	if err != nil {
		os.Exit(10)
	}

	r, err := strconv.Atoi(fields[1])
	if err != nil {
		os.Exit(11)
	}

	k, err := strconv.Atoi(fields[2])
	if err != nil {
		os.Exit(12)
	}

	b, err := strconv.Atoi(fields[3])
	if err != nil {
		os.Exit(12)
	}

	avgRps, peakRps, storBytesByDay := p1.CalculatorBofE(dau, r, k, b)
	fmt.Println("avg_rps=", avgRps)
	fmt.Println("peak_rps=", peakRps)
	fmt.Println("storage_bytes_per=", storBytesByDay)
}
