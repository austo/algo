package main

import (
	"bufio"
	"container/heap"
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
	n := len(g) + 1
	dist := make([]int, n)
	h := minHeap{
		heap: make([]heapNode, n),
		pos:  make([]int, n),
	}

	for i := 0; i < n; i++ {
		dist[i] = MAX
		h.heap[i] = heapNode{i, MAX}
		h.pos[i] = i
	}

	dist[start] = 0
	h.Update(start, 0)

	for h.Len() > 0 {
		hn := heap.Pop(&h).(heapNode)
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

type minHeap struct {
	heap []heapNode
	pos  []int
}

func (h minHeap) Len() int           { return len(h.heap) }
func (h minHeap) Less(i, j int) bool { return h.heap[i].dist < h.heap[j].dist }
func (h minHeap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
	h.pos[h.heap[i].v] = i
	h.pos[h.heap[j].v] = j
}

func (h *minHeap) Push(x interface{}) {
	h.heap = append(h.heap, x.(heapNode))
}

func (h *minHeap) Pop() interface{} {
	n := h.Len() - 1
	x := h.heap[n]
	h.heap = h.heap[0:n]
	return x
}

func (h *minHeap) Update(v, dist int) {
	index := h.pos[v]
	h.heap[index].dist = dist
	heap.Fix(h, index)
}

func (h minHeap) Has(v int) bool {
	return h.pos[v] < h.Len()
}
