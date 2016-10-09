package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`\s+`)

const MAX = 1000000

type edge struct {
	v int
	w int
}

func main() {
	var r io.Reader

	if len(os.Args) > 1 {
		if f, err := os.Open(os.Args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			r = f
			defer f.Close()
		}
	} else {
		r = os.Stdin
	}

	g := make(map[int][]edge, 200)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if err := appendVertexInfo(g, scanner.Text()); err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
	}
	dijkstra(g, 1)
}

func dijkstra(g map[int][]edge, start int) {
	dist := make([]int, len(g)+1)
	h := newHeap(len(g) + 1)

	for i := 0; i <= len(g); i++ {
		dist[i] = MAX
		h.heap[i] = heapNode{i, MAX}
		h.pos[i] = i
	}

	dist[start] = 0
	h.Update(start, 0)

	for len(h.heap) > 0 {
		hn := h.Pop()
		u := hn.v

		if dist[u] == MAX {
			continue
		}

		for _, edge := range g[u] {
			if !h.Has(edge.v) {
				continue
			}
			if dist[u]+edge.w < dist[edge.v] {
				dist[edge.v] = dist[u] + edge.w
				h.Update(edge.v, dist[edge.v])
			}
		}
	}
	printSelected(dist)
}

func printSelected(dist []int) {
	for i, v := range []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197} {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", dist[v])
	}
	fmt.Println()
}

func appendVertexInfo(g map[int][]edge, s string) error {
	parts := re.Split(s, -1)
	if len(parts) < 2 {
		return errors.New("Invalid line")
	}
	i := parts[0]
	v, err := strconv.Atoi(i)
	if err != nil {
		return err
	}
	for _, e := range parts[1:] {
		info := strings.Split(e, ",")
		if len(info) < 2 {
			continue
		}
		v2, err := strconv.Atoi(info[0])
		if err != nil {
			return err
		}
		w, err := strconv.Atoi(info[1])
		if err != nil {
			return err
		}
		g[v] = append(g[v], edge{v2, w})
	}
	return nil
}

type heapNode struct {
	v    int
	dist int
}

type heap struct {
	heap []heapNode
	pos  []int
}

func newHeap(n int) heap {
	h := heap{
		heap: make([]heapNode, n),
		pos:  make([]int, n),
	}
	return h
}

func (h heap) Init() {
	n := len(h.heap)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func (h *heap) Push(v heapNode) {
	h.heap = append(h.heap, v)
	up(*h, len(h.heap)-1)
}

func (h *heap) Pop() heapNode {
	n := len(h.heap) - 1
	h.Swap(0, n)
	down(*h, 0, n)
	x := h.heap[n]
	h.heap = h.heap[0:n]
	return x
}

func (h heap) Update(v, dist int) {
	index := h.pos[v]
	h.heap[index].dist = dist
	down(h, index, len(h.heap))
	up(h, index)
}

func (h heap) Has(v int) bool {
	return h.pos[v] < len(h.heap)
}

func (h heap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
	h.pos[h.heap[i].v] = i
	h.pos[h.heap[j].v] = j
}

func up(h heap, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || h.heap[j].dist >= h.heap[i].dist {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h heap, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.heap[j1].dist >= h.heap[j2].dist {
			j = j2 // = 2*i + 2  // right child
		}
		if h.heap[j].dist >= h.heap[i].dist {
			break
		}
		h.Swap(i, j)
		i = j
	}
}
