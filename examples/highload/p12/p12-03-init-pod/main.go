package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	params := make(map[string]int64)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err != nil {
			continue
		}
		params[key] = val
	}

	initialDelay := params["initialDelaySeconds"]
	period := params["periodSeconds"]
	failureThreshold := params["failureThreshold"]
	successThreshold := params["successThreshold"]
	appStartup := params["appStartupSeconds"]

	// Момент, когда kubelet зафиксирует failureThreshold-ю подряд неуспешную пробу
	detect := initialDelay + (failureThreshold-1)*period

	// Если приложение не успело стартовать к моменту detect — вечный CrashLoop
	if appStartup > detect {
		fmt.Printf("detect=%ds ready=- verdict=crashloop\n", detect)
		return
	}

	// Индекс первой пробы с временем >= appStartupSeconds
	var k int64
	if appStartup <= initialDelay {
		k = 0
	} else {
		diff := appStartup - initialDelay
		k = (diff + period - 1) / period // ceil(diff / period)
	}

	// Готовность наступает после successThreshold подряд успешных проб
	ready := initialDelay + (k+successThreshold-1)*period

	fmt.Printf("detect=%ds ready=%ds verdict=ok\n", detect, ready)
}
