package main

import "fmt"

func main() {
	a := []int{8, 5, 6, -2, 4, 9, 88, 63, 8, 7, 11, -32, -3, 23, 18, 4, 55}
	liss := lis(a)
	fmt.Println(a)
	moves := selsort(a)
	fmt.Println(a)
	fmt.Printf("elements: %d\n", len(a))
	fmt.Printf("moves: %d\n", moves)
	fmt.Printf("lis: %d\n", liss)
}

func selsort(a []int) int {
	moves := 0
	n := len(a)

	for i := 0; i < n-1; i++ {
		min := i // assume min is first element
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
			moves++
		}
	}
	return moves
}

// lis returns the longest increasing subsequence in a
func lis(a []int) int {
	lis := make([]int, 0, len(a))

	// initialize LIS values for all indices
	for range a {
		lis = append(lis, 1)
	}

	// Compute optimized LIS values in bottom up manner
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}

	// Pick maximum of all LIS values
	max := 0
	for _, v := range lis {
		if max < v {
			max = v
		}
	}

	return max
}
