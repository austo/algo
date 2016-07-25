package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	for _, v := range [][]int{
		[]int{6, -6, 8, 4, -12, 9, 8, -8},
		[]int{4, 6, -10, 8, 9, 10, -19, 10, -18, 20, 25},
		[]int{4, 6, -10, 8, 9, 10, -19, 10, -18, 20, 25, -7, -13},
		[]int{4, 6, -10, 8, 9, 10, -19, 10, -18, 20, 25, -7, -13, 1},
	} {
		l := makeList(v)
		// printList(cancelSums2(l))
		fmt.Println(cancelSums(l))
	}
}

func cancelSums(head *node) []int {
	// 1. walk the list and elements to map
	// 2. decrement negative map entry if found
	// 3. when the elements add to zero,
	// 		clear the map to handle mutiple summands
	sum := 0
	m := map[int]int{}
	for np := head; np != nil; np = np.next {
		sum += np.data
		if sum == 0 {
			m = map[int]int{}
		} else {
			if _, ok := m[-np.data]; !ok {
				m[np.data]++
			} else {
				delete(m, -np.data)
			}
		}
	}

	// TODO: some recursive thing to get rid of the remaining sums
	if v, ok := m[sum]; ok {
		m = map[int]int{
			sum: v,
		}
	}

	nodes := []int{}
	for k, v := range m {
		for i := 0; i < v; i++ {
			nodes = append(nodes, k)
		}
	}
	for i := 0; i < len(nodes); i++ {
		s := sum - nodes[i]
		if v, ok := m[s]; ok {
			m = map[int]int{
				s:        v,
				nodes[i]: m[nodes[i]],
			}
		}
	}
	nodes = []int{}
	for k, v := range m {
		for i := 0; i < v; i++ {
			nodes = append(nodes, k)
		}
	}

	return nodes
}

func cancelSums2(head *node) *node {
	stack := []int{}
	m := map[int]*node{}
	sum := 0
	np := head
	for np != nil {
		sum += np.data
		if _, ok := m[sum]; ok {
			for stack[len(stack)-1] != sum {
				top := stack[len(stack)-1]
				delete(m, top)
			}
			m[sum].next = np
		} else if sum == 0 {
			head = np.next
			m = map[int]*node{}
			stack = []int{}
		} else {
			stack = append(stack, sum)
			m[sum] = np
		}
		np = np.next
	}
	return head
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
