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
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
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

	// allocatable
	allocCPU := params["capacityCPU"] - params["kubeReservedCPU"] - params["systemReservedCPU"]
	allocMem := params["capacityMem"] - params["kubeReservedMem"] - params["systemReservedMem"] - params["evictionHardMem"]

	// free = allocatable - daemonset, не меньше 0
	freeCPU := allocCPU - params["daemonsetCPU"]
	if freeCPU < 0 {
		freeCPU = 0
	}
	freeMem := allocMem - params["daemonsetMem"]
	if freeMem < 0 {
		freeMem = 0
	}

	// три предела
	limitCPU := int64(0)
	if params["podRequestCPU"] > 0 {
		limitCPU = freeCPU / params["podRequestCPU"]
	}
	limitMem := int64(0)
	if params["podRequestMem"] > 0 {
		limitMem = freeMem / params["podRequestMem"]
	}
	limitPods := params["maxPods"]

	n := limitCPU
	if limitMem < n {
		n = limitMem
	}
	if limitPods < n {
		n = limitPods
	}

	var bounds []string
	if limitCPU == n {
		bounds = append(bounds, "cpu")
	}
	if limitMem == n {
		bounds = append(bounds, "mem")
	}
	if limitPods == n {
		bounds = append(bounds, "maxpods")
	}

	fmt.Printf("allocatable=%dm/%dMi free=%dm/%dMi pods=%d bound=%s\n",
		allocCPU, allocMem, freeCPU, freeMem, n, strings.Join(bounds, "+"))
}
