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

// var re = regexp.MustCompile(`(^\d+)(?:\s+?(\d+,\d+))*`)
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
	// for k := range graph {
	// 	fmt.Println(k)
	// }
	// fmt.Println(len(graph))
	// for i := 1; i <= 200; i++ {
	// 	fmt.Println(i, graph[i])
	// 	fmt.Println()
	// }
	dijkstra(g, 1)
}

func dijkstra(g map[int][]edge, start int) {
	dist := make([]int, len(g)+1)
	intree := make([]bool, len(g)+1)

	for i := 0; i <= len(g); i++ {
		dist[i] = MAX
	}

	dist[start] = 0

	for i := 1; i <= len(g); i++ {
		u := minDistance(dist, intree)
		intree[u] = true

		for v := 1; v <= len(g); v++ {
			// Update dist[v] only if v is not in intree, there is an
			// edge from u to v, and total weight of path from start to
			// v through u is smaller than current value of dist[v]
			e, ok := hasEdge(g, u, v)
			if !ok || intree[v] || dist[u] == MAX {
				continue
			}
			if dist[u]+e.w < dist[v] {
				dist[v] = dist[u] + e.w
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

func hasEdge(g map[int][]edge, u, v int) (edge, bool) {
	for _, e := range g[u] {
		if e.v == v {
			return e, true
		}
	}
	return edge{}, false
}

func minDistance(dist []int, intree []bool) int {
	min := MAX
	minIndex := -1

	for i, v := range dist {
		if !intree[i] && v <= min {
			min = v
			minIndex = i
		}
	}
	return minIndex
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
