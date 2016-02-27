package main

func bfs(start int, nodes map[int][]int, fn func(int)) {
	frontier := []int{start}
	visited := map[int]bool{}

	for len(frontier) > 0 {
		next := []int{}
		for _, node := range frontier {
			visited[node] = true
			fn(node)
			for _, n := range nodes[node] {
				if !visited[n] {
					next = append(next, n)
				}
			}
		}
		frontier = next
	}
}

func bfsFrontier(node int, nodes map[int][]int, visited map[int]bool) []int {
	next := []int{}
	for _, n := range nodes[node] {
		if !visited[n] {
			next = append(next, n)
		}
	}
	return next
}
