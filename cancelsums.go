package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	l := makeList([]int{6, -6, 8, 4, -12, 9, 8, -8})
	r := cancelSums2(l)
	// printList(r)
	printNodes(r)
}

func cancelSums2(head *node) []*node {
	// walk the list and tally the elements
	// when the elements add to zero, reset the head of the list to the next element
	sum := 0
	nodes := []*node{}
	for np := head; np != nil; np = np.next {
		sum += np.data
		nodes = append(nodes, np)
		if sum == 0 {
			nodes = []*node{}
		}
	}
	return nodes
}

func cancelSums(head *node) *node {
	m := make(map[int]int)
	pos := 0
	neg := 0
	for np := head; np != nil; np = np.next {
		d := np.data
		m[d]++
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

// func find(arr []int, sum int) []int {
// 	currSum := arr[0]
// 	for

// }

func find(k, m map[int]int) (bool, []int) {
	if v, ok := m[k]; !ok {
		return false, nil
	} else {
		return true, []int{k}
	}
	// get smallest number < k
	var j int
	if k < 0 {
		j = -1
	} else {
		j = 1
	}
	rest := []int{}
	for j != k {
		if v, ok := m[j]; !ok {
			if j < 0 {
				j--
			} else {
				j++
			}
		}
	}
	if k < 0 {
		for j := -1; j > k; j-- {
			if 
		}
	} else {
		for j := 1; j < k; j++ {

		}
	}

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

func printNodes(nodes []*node) {
	for i, v := range nodes {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v.data)
	}
	fmt.Println()
}
