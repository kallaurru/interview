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

	grace := params["terminationGracePeriodSeconds"]
	preStop := params["preStopSleepSeconds"]
	appShutdown := params["appShutdownSeconds"]
	endpoint := params["endpointPropagationSeconds"]

	// Потеря трафика: preStop короче, чем распространение endpoint-изменений
	trafficLoss := preStop < endpoint

	var sigtermStr string
	var exit int64
	var killed bool

	if preStop >= grace {
		// preStop не успевает завершиться до истечения grace period —
		// SIGTERM не отправляется, kubelet сразу шлёт SIGKILL.
		sigtermStr = "-"
		exit = grace
		killed = true
	} else {
		// SIGTERM отправляется сразу после preStop.
		sigtermStr = fmt.Sprintf("%ds", preStop)
		natural := preStop + appShutdown
		if natural > grace {
			// Приложение не успело дослужить — SIGKILL по grace period.
			exit = grace
			killed = true
		} else {
			exit = natural
			killed = false
		}
	}

	killedStr := "no"
	if killed {
		killedStr = "yes"
	}
	trafficStr := "no"
	if trafficLoss {
		trafficStr = "yes"
	}

	fmt.Printf("sigterm=%s exit=%ds killed=%s traffic_loss=%s\n",
		sigtermStr, exit, killedStr, trafficStr)
}
