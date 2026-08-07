package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var l, ll, w int
	count := 0
	_, err := fmt.Fscan(r, &count)
	if err != nil {
		os.Exit(1)
	}
	if count <= 0 {
		os.Exit(11)
	}
	for i := 0; i < count; i++ {
		var inL, inLL, inW string
		flg := 0
		_, err = fmt.Fscan(r, &inL, &inLL, &inW)
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		l, err = convert(inL)
		if err != nil {
			fmt.Printf("Ошибка конвертации строки %s\n. Error - %s", inL, err.Error())
			os.Exit(3)
		}
		ll, err = convert(inLL)
		if err != nil {
			fmt.Printf("Ошибка конвертации строки %s\n. Error - %s", inLL, err.Error())
			os.Exit(3)
		}
		w, err = convert(inW)
		if err != nil {
			fmt.Printf("Ошибка конвертации строки %s\n. Error - %s", inW, err.Error())
			os.Exit(3)
		}
		l, ll, w = littleCalc(l, ll, w, &flg)
		printer(flg, l, ll, w)
	}
}

func littleCalc(l, lambda, w int, flg *int) (int, int, int) {
	if l == -1 {
		l = (lambda * w) / 1000
		*flg = 1
	}

	if lambda == -1 {
		lambda = (l * 1000) / w
		*flg = 2
	}

	if w == -1 {
		w = (l * 1000) / lambda
		*flg = 3
	}

	return l, lambda, w
}

func convert(in string) (int, error) {
	if in == "?" {
		return -1, nil
	}

	return strconv.Atoi(in)
}

func printer(flg, l, lambda, w int) {
	switch flg {
	default:
		fmt.Printf("%d\n", flg)
	case 1:
		fmt.Printf("%d\n", l)
	case 2:
		fmt.Printf("%d\n", lambda)
	case 3:
		fmt.Printf("%d\n", w)
	}
}
