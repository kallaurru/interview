package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Task struct {
	PoolID int
	Start  int
	End    int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var w, n int

	_, err := fmt.Fscan(r, &w, &n)
	if err != nil {
		os.Exit(1)
	}

	alltime := 0
	makespan := 0
	mngr := make([]Task, w)
	for i := 0; i < n; i++ {
		var val int
		_, err = fmt.Fscan(r, &val)
		if err != nil {
			os.Exit(2)
		}
		alltime += val
		if i < w {
			mngr[i] = Task{PoolID: i, Start: 0, End: val}
			if val > makespan {
				makespan = val
			}
			continue
		}
		// работаем с освободившимися воркерами
		sort.Slice(mngr, func(i, j int) bool {
			if mngr[i].End == mngr[j].End {
				return mngr[i].PoolID < mngr[j].PoolID
			}
			return mngr[i].End < mngr[j].End // tie-break по метке
		})
		if len(mngr) == 0 {
			continue
		}

		mngr[0] = Task{PoolID: i % w, Start: mngr[0].End, End: mngr[0].End + val}
		if mngr[0].End > makespan {
			makespan = mngr[0].End
		}
	}

	idle := makespan*w - alltime

	fmt.Printf("%d %d\n", makespan, idle)
}
