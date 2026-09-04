package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	opAcquire = "acquire"
)

type PoolMngr struct {
	dials  int
	closes int
	idle   int
	busy   int
}

func New() PoolMngr {
	return PoolMngr{}
}

func (pm PoolMngr) Manage(op string, maxIdle, maxOpen int) PoolMngr {
	if op == opAcquire {
		if pm.idle == 0 {
			pm.dials++
			pm.busy++
			return pm
		}
		if maxOpen-pm.busy <= 0 {
			return pm
		}
		// можем достать из пула?
		if pm.idle > 0 {
			pm.idle--
		} else {
			pm.dials++
		}
		pm.busy++
		return pm
	}
	// op release
	pm.busy--
	if pm.idle < maxIdle {
		pm.idle++
	} else {
		pm.closes++
	}
	return pm
}

func (pm PoolMngr) Dials() int {
	return pm.dials
}

func (pm PoolMngr) Closes() int {
	return pm.closes
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		n, maxOpen, maxIdle int
	)

	_, err := fmt.Fscan(r, &maxOpen, &maxIdle, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	// todo
	/*	items := []string{"acquire",
		"acquire",
		"release",
		"acquire",
		"release",
		"release"}
	*/

	pm := New()
	for i := 0; i < n; i++ {
		var (
			op string
		)
		_, err = fmt.Fscan(r, &op)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}

		// todo
		//op = items[i]

		pm = pm.Manage(op, maxIdle, maxOpen)
	}
	fmt.Printf("%d %d\n", pm.Dials(), pm.Closes())
}
