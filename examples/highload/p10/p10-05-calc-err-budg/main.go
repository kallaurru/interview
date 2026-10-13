package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	// Read SLO
	if !sc.Scan() {
		return
	}
	slo, err := strconv.ParseFloat(strings.TrimSpace(sc.Text()), 64)
	if err != nil {
		return
	}

	// Read good and total
	if !sc.Scan() {
		return
	}
	parts := strings.Fields(sc.Text())
	if len(parts) < 2 {
		return
	}
	good, err1 := strconv.ParseInt(parts[0], 10, 64)
	total, err2 := strconv.ParseInt(parts[1], 10, 64)
	if err1 != nil || err2 != nil {
		return
	}

	bad := total - good
	allowed := (1.0 - slo/100.0) * float64(total)
	remaining := allowed - float64(bad)

	status := "OK"
	if float64(bad) > allowed {
		status = "BURNED"
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintf(out, "allowed %.2f\n", allowed)
	fmt.Fprintf(out, "remaining %.2f\n", remaining)
	fmt.Fprintf(out, "status %s\n", status)
}
