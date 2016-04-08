package main

import "fmt"

var graph = map[int][]int{
	1: []int{2, 3, 4},
	2: []int{3},
	3: []int{2},
	4: []int{1},
	5: []int{},
}

func main() {
	result := findConnectedChildNodes(1, graph)
	prettyPrint(result)
}

func findConnectedChildNodes(start int, nodes map[int][]int) [][]int {
	result := [][]int{}
	visited := map[int]bool{}
	for _, n := range graph[start] {
		visited[n] = true
		group := map[int]bool{}
		group[n] = true
		for _, nn := range graph[n] {
			if !visited[nn] {
				group[nn] = true
			}
		}
		if len(group) > 1 {
			list := []int{}
			for k, _ := range group {
				if k != start {
					list = append(list, k)
				}
			}
			result = append(result, list)
		}
	}
	return result
}

func prettyPrint(vv [][]int) {
	for _, v := range vv {
		for i, k := range v {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(k)
		}
		fmt.Println()
	}
}
