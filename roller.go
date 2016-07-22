package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// n := 9
	// k := 3
	// breaks := []int{1, 2, 5}

	n := 12
	k := 2
	breaks := []int{3, 5, 7, 9}

	// fmt.Println(checkSpaces([]int{1, 2, 3}, 12, 2))
	// fmt.Println(cache)

	go func() {
		minSeen := n
		var winner []int
		for i := 0; i < n; i++ {
			nSeen, seen := expand([]int{i}, n, k, map[int]bool{}, breaks)
			if nSeen < minSeen {
				minSeen = nSeen
				winner = seen
			}
		}
		fmt.Printf("min covered: %d %v\n", minSeen, winner)
	}()

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGINT)
	time.Sleep(20 * time.Millisecond)
	// fmt.Println(cache)
	os.Exit(1)

}

func expand(path []int, n, k int, visited map[int]bool, breaks []int) (int, []int) {
	fmt.Printf("path: %v\n", path)
	visited[path[len(path)-1]] = true

	seen := checkSpaces(path, n, k)
	fmt.Printf("seen: %v\n", seen)

	if containsAll(seen, breaks) {
		fmt.Println("seen contains all breaks")
		return len(seen), seen
	}

	minSeen := n
	var winner []int
	for i := 0; i < n; i++ {
		if !visited[i] {
			p := make([]int, len(path), len(path)+1)
			copy(p, path)
			p = append(p, i)
			remainingSeen, candidate := expand(p, n, k, cloneMap(visited), breaks)
			if remainingSeen < minSeen {
				minSeen = remainingSeen
				winner = candidate
			}
		}
	}
	fmt.Printf("min covered: %d %v (path: %v)\n", minSeen, winner, path)

	return minSeen, winner
}

var (
	cache = map[string][]int{}
)

func checkSpaces(path []int, n, k int) []int {
	if len(path) == 0 {
		return []int{}
	}
	if len(path) == 1 {
		key := strconv.Itoa(path[0])
		if seen, ok := cache[key]; ok {
			return seen
		} else {
			for i := path[0]; i < path[0]+k && i < n; i++ {
				seen = append(seen, i)
			}
			cache[key] = seen
			return seen
		}
	}
	key := cacheKey(path)
	if seen, ok := cache[key]; ok {
		return seen
	}
	prevSeen := checkSpaces(path[:len(path)-1], n, k)
	thisSeen := checkSpaces([]int{path[len(path)-1]}, n, k)
	nowSeen := union(prevSeen, thisSeen)
	cache[key] = nowSeen
	return nowSeen
}

func cloneMap(m map[int]bool) map[int]bool {
	cloned := map[int]bool{}
	for k, v := range m {
		cloned[k] = v
	}
	return cloned
}

func union(a, b []int) []int {
	m := map[int]struct{}{}
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		m[v] = struct{}{}
	}

	r := make([]int, 0, len(m))
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}

func containsAll(a, b []int) bool {
	m := map[int]struct{}{}
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; !ok {
			return false
		}
	}
	return true
}

func cacheKey(a []int) string {
	var s string
	for i, v := range a {
		if i > 0 {
			s += "-"
		}
		s += strconv.Itoa(v)
	}
	return s
}
