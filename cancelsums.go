package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	l := makeList([]int{6, -6, 8, 4, -12, 9, 8, -8})
	r := cancelSums(l)
	printList(r)
}

func cancelSums(head *node) *node {
	m := make(map[int]struct{})
	pos := 0
	neg := 0
	for np := head; np != nil; np = np.next {
		d := np.data
		m[d] = struct{}{}
		if d > 0 {
			pos += d
		} else if d < 0 {
			neg += d
		}
	}
	if pos == neg {
		return nil
	}
	return &node{data: pos + neg}
}

func makeList(ints []int) *node {
	n := len(ints)
	if n < 1 {
		return nil
	}
	head := &node{data: ints[0]}
	np := head
	for i := 1; i < n; i++ {
		next := &node{data: ints[i]}
		np.next = next
		np = np.next
	}
	return head
}

func printList(head *node) {
	i := 0
	for np := head; np != nil; np = np.next {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(np.data)
		i++
	}
	fmt.Println()
}
