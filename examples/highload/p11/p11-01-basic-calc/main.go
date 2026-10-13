package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
DAU — пользователей в день,
R — среднее число запросов на пользователя в день,
B — средний размер записи в байтах,
P — пиковый коэффициент (peak factor),
REPL — фактор репликации,
DAYS — число дней хранения.
*/
type Calculator struct {
	ceil func(a, b int64) int64
	repl int64
	days int64
	dau  int64
	r    int64
	b    int64
	p    int64
}

/*
avg_qps=<x>
peak_qps=<y>
storage_gb=<z>
bandwidth_mbps=<w>
- x = ceil(DAU·R / 86400)
- y = avg_qps · P
- z = ceil(DAU·R·B·DAYS·REPL / 1000000000)
- w = ceil(avg_qps·B·8 / 1000000)
*/
type Result struct {
	AvgQ   int64
	PeakQ  int64
	StorGb int64
	Bw     int64
}

func New(dau, days, repl, r, b, p int64) Calculator {
	ceil := func(a, b int64) int64 {
		return (a + b - 1) / b
	}

	return Calculator{
		ceil: ceil,
		dau:  dau,
		days: days,
		repl: repl,
		r:    r,
		b:    b,
		p:    p,
	}
}
func (c Calculator) Calc() Result {
	const ts int64 = 86400
	const storDiv int64 = 1000000000
	const bwDiv int64 = 1000000
	const bwItem int64 = 8

	result := Result{}

	result.AvgQ = c.ceil(c.dau*c.r, ts)
	result.PeakQ = result.AvgQ * c.p
	result.StorGb = c.ceil(c.dau*c.r*c.b*c.days*c.repl, storDiv)
	result.Bw = c.ceil(result.AvgQ*c.b*bwItem, bwDiv)

	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var dau, r, b, p, repl, days int64

	fmt.Fscan(in, &dau, &r, &b, &p, &repl, &days)
	calc := New(dau, days, repl, r, b, p)

	result := calc.Calc()

	fmt.Printf("avg_qps=%d\n", result.AvgQ)
	fmt.Printf("peak_qps=%d\n", result.PeakQ)
	fmt.Printf("storage_gb=%d\n", result.StorGb)
	fmt.Printf("bandwidth_mbps=%d\n", result.Bw)
}
