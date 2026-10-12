package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !scanner.Scan() {
		return
	}
	slo, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	if err != nil || slo >= 1000 || slo < 0 {
		return
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		good, e1 := strconv.ParseInt(fields[0], 10, 64)
		total, e2 := strconv.ParseInt(fields[1], 10, 64)
		if e1 != nil || e2 != nil || total <= 0 || good < 0 || good > total {
			continue
		}

		errCount := total - good
		denom := total * (1000 - slo)
		burnX100 := errCount * 100000 / denom

		whole := burnX100 / 100
		frac := burnX100 % 100

		var status string
		switch {
		case burnX100 >= 1440:
			status = "PAGE"
		case burnX100 >= 100:
			status = "TICKET"
		default:
			status = "OK"
		}

		fmt.Fprintf(out, "%d.%02d %s\n", whole, frac, status)
	}
}
