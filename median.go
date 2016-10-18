package main

import "fmt"

/*
 * The goal of this problem is to implement the "Median Maintenance" algorithm (covered in the Week 5 lecture on heap applications).
 * The text file contains a list of the integers from 1 to 10000 in unsorted order; you should treat this as a stream of numbers, arriving one by one.
 * Letting xi denote the ith number of the file, the kth median mk is defined as the median of the numbers x1,…,xk.
 * (So, if k is odd, then mk is ((k+1)/2)th smallest number among x1,…,xk; if k is even, then mk is the (k/2)th smallest number among x1,…,xk.)

 * In the box below you should type the sum of these 10000 medians, modulo 10000 (i.e., only the last 4 digits). That is, you should compute (m1+m2+m3+⋯+m10000)mod10000.
 * OPTIONAL EXERCISE: Compare the performance achieved by heap-based and search-tree-based implementations of the algorithm.
 */

func main() {
	hmin := make(maxHeap, 0, 10000)
	hmax := make(minHeap, 0, 10000)

	var next int

	var sum int64

	for {
		_, err := fmt.Scanf("%d\n", &next)
		if err != nil {
			break
		}
		minlen, maxlen := len(hmin), len(hmax)
		if minlen == 0 && maxlen == 0 {
			sum += int64(next)
			hmax.Push(next)
			continue
		}

		if minlen < maxlen {
			if maxlen-minlen > 1 {
				panic("unbalanced: min")
			}
			high := hmax[0]
			if next < high {
				hmin.Push(next)
			} else {
				hmin.Push(hmax.Pop())
				hmax.Push(next)
			}
			if len(hmin) != len(hmax) {
				panic("still unbalanced: min")
			}
			sum += int64(hmin[0])
		} else if maxlen < minlen {
			if minlen-maxlen > 1 {
				panic("unbalanced: max")
			}
			low := hmin[0]
			if next >= low {
				hmax.Push(next)
			} else {
				hmax.Push(hmin.Pop())
				hmin.Push(next)
			}
			if len(hmin) != len(hmax) {
				panic("still unbalanced: max")
			}
			sum += int64(hmin[0])
		} else {
			if next >= hmin[0] {
				hmax.Push(next)
				sum += int64(hmax[0])
			} else {
				hmin.Push(next)
				sum += int64(hmin[0])
			}
		}
	}

	fmt.Printf("minlen = %d, maxlen = %d\n", len(hmin), len(hmax))
	fmt.Printf("sum = %d, last 4 digits = %d\n", sum, sum%10000)
}

type minHeap []int

func (h *minHeap) Push(v int) {
	push((*[]int)(h), v, func(a []int, i, j int) bool {
		return a[i] < a[j]
	})
}

func (h *minHeap) Pop() int {
	return pop((*[]int)(h), func(a []int, i, j int) bool {
		return a[i] < a[j]
	})
}

type maxHeap []int

func (h *maxHeap) Push(v int) {
	push((*[]int)(h), v, func(a []int, i, j int) bool {
		return a[i] > a[j]
	})
}

func (h *maxHeap) Pop() int {
	return pop((*[]int)(h), func(a []int, i, j int) bool {
		return a[i] > a[j]
	})
}

type pred func(h []int, i, j int) bool

func push(h *[]int, v int, ordered pred) {
	*h = append(*h, v)
	up(*h, ordered)
}

func pop(h *[]int, ordered pred) int {
	old := *h
	n := len(old) - 1
	old[0], old[n] = old[n], old[0]
	down(old, 0, n, ordered)
	x := old[n]
	*h = old[0:n]
	return x
}

func heapify(h []int, ordered pred) {
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n, ordered)
	}
}

func up(h []int, ordered pred) {
	j := len(h) - 1
	for {
		i := (j - 1) / 2 // parent
		if i == j || !ordered(h, j, i) {
			break
		}
		h[i], h[j] = h[j], h[i]
		j = i
	}
}

func down(h []int, i, n int, ordered pred) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // guard against overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && !ordered(h, j1, j2) {
			j = j2 // 2*i + 2 // right child
		}
		if !ordered(h, j, i) {
			break
		}
		h[i], h[j] = h[j], h[i]
		i = j
	}
}
