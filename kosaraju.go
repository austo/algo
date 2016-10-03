package main

import (
	"fmt"
	"io"
	"os"
)

var graph = make(map[int][]int, 1000000)
var rgraph = make(map[int][]int, 1000000)
var largestVertex int = 0

func main() {
	var r io.Reader

	if len(os.Args) > 1 {
		if f, err := os.Open(os.Args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			r = f
		}
	} else {
		r = os.Stderr
	}
	var from, to int
	for {
		_, err := fmt.Fscanln(r, &from, &to)
		if err != nil {
			break
		}
		graph[from] = append(graph[from], to)
		rgraph[to] = append(rgraph[to], from)
		if from > largestVertex {
			largestVertex = from
		}
		if to > largestVertex {
			largestVertex = to
		}
	}
	fmt.Println(largestVertex)
	finished := getFinishingTimes(rgraph)
	fmt.Println(len(finished))
	sccs := findSccs(graph, finished)
	fmt.Println(len(sccs))
	topSccs := topNSccs(sccs, 5)
	fmt.Println(topSccs)

	for i := len(topSccs) - 1; i >= 0; i-- {
		fmt.Printf("%d", topSccs[i].size)
		if i > 0 {
			os.Stdout.WriteString(",")
		}
	}
	fmt.Println()
}

type scount struct {
	parent int
	size   int
}

func topNSccs(sccs map[int]int, n int) []scount {
	rv := make([]scount, n)
	for k, v := range sccs {
		if v <= rv[0].size {
			continue
		}
		// "online" insertion sort
		i := n - 1
		for i > 0 && v <= rv[i].size {
			i--
		}
		copy(rv[0:i], rv[1:i+1])
		rv[i] = scount{k, v}
	}
	return rv
}

func findSccs(g map[int][]int, finished []int) map[int]int {
	visited := make(map[int]bool, len(g))
	sccs := map[int]int{}
	var s int

	var dfs func(int)

	dfs = func(v int) {
		visited[v] = true
		sccs[s] = sccs[s] + 1 // can you do ++ if there's no key??
		for _, n := range g[v] {
			if !visited[n] {
				dfs(n)
			}
		}
	}

	for j := len(finished) - 1; j >= 0; j-- {
		i := finished[j]
		if visited[i] {
			continue
		}
		s = i
		dfs(i)
	}
	return sccs
}

func getFinishingTimes(g map[int][]int) []int {
	finished := make([]int, 0, largestVertex)
	visited := make(map[int]bool, len(g))

	// t := 0

	var dfs func(int)

	dfs = func(v int) {
		visited[v] = true
		for _, n := range g[v] {
			if !visited[n] {
				dfs(n)
			}
		}
		// t++
		finished = append(finished, v)
	}

	for i := largestVertex; i > 0; i-- {
		// for i := 1; i <= largestVertex; i++ {
		if visited[i] {
			continue
		}
		dfs(i)
	}
	return finished
}

func sort(rv []scount) {
	for i := 0; i < len(rv); i++ {
		j := i
		for j > 0 && rv[j].size < rv[j-1].size {
			j--
		}
		sc := rv[i]
		for k := i; k > j; k-- {
			rv[k] = rv[k-1]
		}
		rv[j] = sc
	}
}
