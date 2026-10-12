package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const tolerance = 0.1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	// Первая строка: currentReplicas targetUtilization minReplicas maxReplicas
	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) != 4 {
		return
	}
	currentReplicas, _ := strconv.Atoi(parts[0])
	targetUtil, _ := strconv.Atoi(parts[1])
	minReplicas, _ := strconv.Atoi(parts[2])
	maxReplicas, _ := strconv.Atoi(parts[3])

	// Вторая строка: n
	if !scanner.Scan() {
		return
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return
	}

	var usageSum, requestSum int64
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 {
			continue
		}
		u, err1 := strconv.ParseInt(fields[1], 10, 64)
		r, err2 := strconv.ParseInt(fields[2], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		usageSum += u
		requestSum += r
	}

	if requestSum == 0 {
		// Не должно происходить по условию, но подстрахуемся
		fmt.Printf("avg=0%% desired=%d action=hold\n", currentReplicas)
		return
	}

	// Утилизация по неокруглённому значению
	util := 100.0 * float64(usageSum) / float64(requestSum)

	// Округление «половина вверх» (для положительных math.Round = half away from zero = half up)
	avgRounded := int64(math.Round(util))

	// Отношение метрики к цели
	ratio := util / float64(targetUtil)

	desired := currentReplicas
	action := "hold"

	if math.Abs(ratio-1.0) > tolerance {
		// desired = ceil(currentReplicas * ratio)
		desiredF := float64(currentReplicas) * ratio
		desired = int(math.Ceil(desiredF))

		// Ограничение [minReplicas, maxReplicas]
		if desired < minReplicas {
			desired = minReplicas
		}
		if desired > maxReplicas {
			desired = maxReplicas
		}

		switch {
		case desired > currentReplicas:
			action = "scale_up"
		case desired < currentReplicas:
			action = "scale_down"
		default:
			action = "hold"
		}
	} else {
		// В пределах порога нечувствительности — ничего не делаем
		if desired < minReplicas {
			desired = minReplicas
		}
		if desired > maxReplicas {
			desired = maxReplicas
		}
	}

	fmt.Printf("avg=%d%% desired=%d action=%s\n", avgRounded, desired, action)
}
