package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dispatcher struct {
	x, y int
	buf  int
	c, r int // Объем буфера и кол-во сообщений, которые читает консюмер за тик
}

func New(tics []int, c, r int) Dispatcher {
	d := Dispatcher{c: c, r: r}
	for _, tick := range tics {
		d.buf += tick
		if d.buf > d.c {
			drop := d.buf - d.c
			d.buf -= drop
			d.y += drop
		}
		if d.buf > d.x {
			d.x = d.buf
		}
		// читаем из буф
		left := min(d.r, d.buf)
		// остаток в буфере
		d.buf -= left
	}
	return d
}

func (d *Dispatcher) Stat() (int, int, int) {
	return d.x, d.y, d.buf
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		c, rc, t int
	)

	_, err := fmt.Fscan(r, &c, &rc, &t)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	/*	_, err = fmt.Fscan(r, &n)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
	*/
	var values []int
	var val int
	for i := 0; i < t; i++ {

		_, err = fmt.Fscan(r, &val)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		values = append(values, val)
	}
	d := New(values, c, rc)
	// max depth, dropped, rem in buf last tick
	x, y, z := d.Stat()

	fmt.Printf("max_depth=%d\n", x)
	fmt.Printf("dropped=%d\n", y)
	fmt.Printf("left=%d\n", z)
}
