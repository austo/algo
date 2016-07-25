package main

import (
	"fmt"
	"sort"
)

// func getSums(arr []int) []int {
// 	sums := make([]int, 0, len(arr))
// 	copy(sums, arr)
// 	sort.Ints(sums)

// }

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := getSums(arr)
	fmt.Println(s)
}

func getSums(arr []int) []int {
	m := make(map[int]struct{})
	sums(arr, 0, 0, m)
	r := make([]int, 0, len(m))
	for k, _ := range m {
		r = append(r, k)
	}
	sort.Ints(r)
	return r
}

func sums(arr []int, start, sum int, m map[int]struct{}) {
	n := len(arr)
	if n == start {
		return
	}
	val := sum + arr[start]
	m[val] = struct{}{}
	sums(arr, start+1, val, m)
	sums(arr, start+1, sum, m)
}

// func printList(head *node) {
// 	i := 0
// 	for np := head; np != nil; np = np.next {
// 		if i > 0 {
// 			fmt.Print(", ")
// 		}
// 		fmt.Print(np.data)
// 		i++
// 	}
// 	fmt.Println()
// }

// func printNodes(nodes []*node) {
// 	for i, v := range nodes {
// 		if i > 0 {
// 			fmt.Print(", ")
// 		}
// 		fmt.Print(v.data)
// 	}
// 	fmt.Println()
// }
