package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	var n int
	const maxMoment = 1000001
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(1)
	}

	mxFinish, canceled, errMoment := 0, 0, maxMoment
	var errName string
	moments := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var (
			name, status string
			finish       int
		)

		_, err = fmt.Fscan(r, &name, &status, &finish)
		if err != nil {
			os.Exit(22)
		}
		moments = append(moments, finish)

		if mxFinish < finish {
			mxFinish = finish
		}

		if errMoment > finish && status == "err" {
			errMoment = finish
			errName = name
		}
	}
	if errMoment == maxMoment {
		fmt.Printf("nil %d\n", mxFinish)
		fmt.Printf("canceled %d\n", canceled)
		os.Exit(0)
	}

	slices.Sort(moments)
	idx, found := slices.BinarySearch(moments, errMoment+1)
	if found {
		canceled = len(moments) - idx
	} else {
		if idx < len(moments)-1 {
			canceled = len(moments) - idx
		} else {
			canceled = 0
		}
	}

	fmt.Printf("err %s %d\n", errName, errMoment)
	fmt.Printf("canceled %d\n", canceled)
}
