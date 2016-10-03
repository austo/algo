package main

import (
	"flag"
	"fmt"
	"time"
)

func fib() func(int) int {
	m := map[int]int{}
	var f func(int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}
		f1, f2 := cache(n-1, m, f), cache(n-2, m, f)
		m[n] = f1 + f2
		return f1 + f2
	}
	return f
}

func cache(n int, m map[int]int, f func(int) int) int {
	var f1 int
	if v, ok := m[n]; ok {
		f1 = v
	} else {
		f1 = f(n)
		m[n] = f1
	}
	return f1
}

func fib2(n int) int {
	a := [2]int{0, 1}

	if n < 2 {
		return n
	}

	var f int
	for i := 2; i <= n; i++ {
		f = a[0] + a[1]
		a[0], a[1] = a[1], f
	}
	return f
}

func init() {
	flag.IntVar(&lim, "n", 20, "fib limit")
}

var (
	lim int
)

func main() {
	flag.Parse()
	f := fib()
	start := time.Now()
	f1 := f(lim)
	d := time.Since(start)
	fmt.Printf("%d (%v)\n", f1, d)
	start = time.Now()
	f2 := fib2(lim)
	d = time.Since(start)
	fmt.Printf("%d (%v)\n", f2, d)
}
