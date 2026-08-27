package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	stateNoDeg   = "NO_DEGRADATION"
	stateDisable = "DISABLE"
	stateServ    = "SERVING"
	stateOver    = "OVERLOAD"
)

type Feature struct {
	Name     string
	Priority int
	Cost     int
}

type LoadRegister struct {
	lim      int
	load     int
	features []Feature
}

func New(lim, n int) *LoadRegister {
	return &LoadRegister{
		lim:      lim,
		load:     0,
		features: make([]Feature, 0, n),
	}
}

func (lr *LoadRegister) Collect(name string, priority, cost int) {
	lr.load += cost
	feat := Feature{
		Name:     name,
		Priority: priority,
		Cost:     cost,
	}
	lr.features = append(lr.features, feat)
}

func (lr *LoadRegister) Serve() {
	sort.Slice(lr.features, func(i, j int) bool {
		if lr.features[i].Priority != lr.features[j].Priority {
			return lr.features[i].Priority < lr.features[j].Priority
		}
		return lr.features[i].Name < lr.features[j].Name
	})
	if lr.load <= lr.lim {
		lr.Print(stateNoDeg, "", lr.load)
		return
	}
	for _, feature := range lr.features {
		if feature.Priority >= 100 {
			continue
		}
		lr.load -= feature.Cost
		lr.Print(stateDisable, feature.Name, lr.load)
		if lr.load > lr.lim {
			continue
		}
		break
	}
	if lr.load <= lr.lim {
		lr.Print(stateServ, "", lr.load)
		return
	}
	lr.Print(stateOver, "", lr.load)
}

func (lr *LoadRegister) Print(state, name string, load int) {
	printer(state, name, load)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		c, n int
	)

	_, err := fmt.Fscan(r, &c)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	lr := New(c, n)
	for i := 0; i < n; i++ {
		var (
			name           string
			priority, cost int
		)

		_, err = fmt.Fscan(r, &name, &priority, &cost)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		lr.Collect(name, priority, cost)
	}
	lr.Serve()
}

func printer(state, name string, load int) {
	if state == stateDisable {
		fmt.Printf("%s %s %d\n", state, name, load)
		return
	}
	fmt.Printf("%s %d\n", state, load)
}
