package main

import (
	"fmt"
	"sort"
)

type withDigraph []string

func (v withDigraph) Len() int {
	return len(v)
}

func (v withDigraph) Less(i, j int) bool {
	return compare(v[i], v[j])
}

func (v withDigraph) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

var digraphs = map[string]int{
	"aa": 256,
	"ch": 257,
	"cz": 258,
}

func main() {
	strings := []string{
		"cheese", "cracker", "casserole",
		"crumpet", "croissant", "cookiedough",
		"cookie", "reach", "react", "absolute",
		"aardvark",
	}
	sort.Sort(withDigraph(strings))
	fmt.Println(strings)
}

// compare considers digraphs (as defined in digraphs)
// and has a running time O(n), where n = min(len(a), len(b))
func compare(a, b string) bool {
	la, lb := len(a), len(b)
	if la == 0 && lb == 0 {
		return false
	}
	if lb == 0 {
		return false
	}
	if la == 0 {
		return true
	}
	i, j, wa, wb := 0, 0, 0, 0
	for wa == wb && i < la && j < lb {
		wa, i = consume(a, i)
		wb, j = consume(b, j)
	}
	if wa == wb {
		return la < lb
	}
	return wa <= wb
}

func consume(s string, i int) (int, int) {
	w, ok := isDigraph(s[i:])
	if ok {
		i += 2
	} else {
		w = int(s[i])
		i++
	}
	return w, i
}

func isDigraph(s string) (int, bool) {
	if len(s) < 2 {
		return 0, false
	}
	w, ok := digraphs[s[:2]]
	return w, ok
}
