package main

import "fmt"

// implement a min-heap for integers

type IntHeap []int

func (h IntHeap) Init() {
	// heapify
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func (h *IntHeap) Push(v int) {
	*h = append(*h, v)
	up(*h, len(*h)-1)
}

func (h *IntHeap) Pop() int {
	old := *h
	n := len(old) - 1
	old[0], old[n] = old[n], old[0]
	down(old, 0, n)
	x := old[n]
	*h = old[0:n]
	return x
}

func up(h IntHeap, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || h[j] >= h[i] {
			break
		}
		h[i], h[j] = h[j], h[i]
		j = i
	}
}

func down(h IntHeap, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h[j1] >= h[j2] {
			j = j2 // = 2*i + 2  // right child
		}
		if h[j] >= h[i] {
			break
		}
		h[i], h[j] = h[j], h[i]
		i = j
	}
}

func main() {
	h := IntHeap{}
	for _, v := range []int{2, 1, 5, 3} {
		h.Push(v)
	}
	fmt.Printf("minimum: %d\n", h[0])
	for len(h) > 0 {
		fmt.Printf("%d ", h.Pop())
	}

	fmt.Println()

	h2 := IntHeap{2, 1, 5, 3}
	h2.Init()
	fmt.Printf("minimum: %d\n", h2[0])
	for len(h2) > 0 {
		fmt.Printf("%d ", h2.Pop())
	}

	fmt.Println()
}
