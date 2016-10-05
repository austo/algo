package main

import "fmt"

func hasCycle(g map[int][]int) bool {
	visited := make(map[int]bool, len(g))
	instack := map[int]bool{}

	var isCycle func(int) bool

	isCycle = func(v int) bool {
		visited[v] = true
		instack[v] = true

		for _, i := range g[v] {
			if !visited[i] && isCycle(i) {
				return true
			}
			if instack[i] {
				return true
			}
		}
		delete(instack, v)
		return false
	}

	cycleDetected := false

	for i := range g {
		if !cycleDetected && !visited[i] {
			cycleDetected = isCycle(i)
		}
	}

	return cycleDetected
}

var graph map[int][]int = map[int][]int{
	1: []int{2, 4},
	2: []int{3, 4},
	3: []int{4, 1},
}

func main() {
	fmt.Println(hasCycle(graph))
}
