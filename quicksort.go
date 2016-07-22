package main

import "fmt"

func quicksort(a []int) {
	if len(a) < 2 {
		return
	}

	l, r := 0, len(a)-1

	// pick middle element as the pivot (random will also work,
	// though for testing it may be better to chose the first
	// or last to enable exploration of pathological O(n^2) performance)
	// http://stackoverflow.com/questions/164163/quicksort-choosing-the-pivot

	p := len(a) / 2

	// move pivot to rightmost position
	a[p], a[r] = a[r], a[p]

	// move all elements smaller than pivot to left of pivot
	for i := range a {
		if a[i] < a[r] {
			a[i], a[l] = a[l], a[i]
			l++
		}
	}

	// move pivot to the right of the last smaller element
	a[l], a[r] = a[r], a[l]

	// recurse
	quicksort(a[:l])
	quicksort(a[l+1:])
}

func main() {
	nums := []int{45, 25, 48, -96, 1, 3, -3, -69, 589,
		479, 56, 54, 852, -258, 88, -395, 42, 89, 313, -8,
		-42, 16, 39, 100, 4, 2, 88, 13, -296, 14, 869, 11}
	quicksort(nums)
	fmt.Println(nums)
}
