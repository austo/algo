package main

import "fmt"

func main() {
	// n := 9
	// k := 3
	// breaks := []int{1, 2, 5}

	n := 12
	k := 2
	breaks := []int{3, 5, 7, 9}

	fmt.Println(expand(breaks, n, k))
}

func expand(breaks []int, n, k int) int {
	min := len(breaks) * k
	r := []int{}
	for i := 0; i < k; i++ {
		r = append(r, i)
	}
	o := Odometer(r, len(breaks))

	for {
		v := o.Next()
		if v == nil {
			break
		}
		if ok, pos := legal(n, k, v, breaks); ok {
			c := covered(k, pos...)
			if c < min {
				min = c
			}
		}
	}
	return min
}

func covered(k int, positions ...int) int {
	m := map[int]struct{}{}
	for _, v := range positions {
		for i := 0; i < k; i++ {
			m[v+i] = struct{}{}
		}
	}
	return len(m)
}

func legal(n, k int, v, breaks []int) (bool, []int) {
	pos := []int{}
	for i, p := range v {
		b := breaks[i]
		np := b - p
		if np < 0 || np >= n {
			return false, pos
		}
		pos = append(pos, np)
	}
	return true, pos
}

/* Supporting "package" */

// O is an "odometer" type
type O struct {
	limit   int
	curr    int
	r       []int
	clock   []int
	started bool
}

// Odometer implements a rolling odometer with a given range
// see python's itertools.product(*iterables[, repeat])
// https://docs.python.org/2.7/library/itertools.html#itertools.product
func Odometer(r []int, repeat int) *O {
	o := &O{
		r:     r,
		limit: len(r),
		clock: make([]int, repeat),
		curr:  repeat - 1,
	}
	return o
}

// Next gets the subsequent state of the odometer
func (o *O) Next() []int {
	if len(o.r) == 0 {
		return nil
	}
	if o.started {
		if o.clock[o.curr] == o.limit-1 {
			for o.curr >= 0 && o.clock[o.curr] == o.limit-1 {
				o.curr--
			}
			if o.curr < 0 {
				return nil
			}
			o.clock[o.curr]++
			for i := o.curr + 1; i < len(o.clock); i++ {
				o.clock[i] = 0
			}
			o.curr = len(o.clock) - 1
		} else {
			o.clock[o.curr]++
		}
	} else {
		o.started = true
	}
	result := make([]int, 0, len(o.clock))
	for _, v := range o.clock {
		result = append(result, o.r[v])
	}
	return result
}
