package main

import "fmt"

func mergesort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	// create two lists, one for even-indexed elements
	// and one for odd-indexed elements.

	left, right := []int{}, []int{}

	for i, v := range a {
		if i%2 == 0 {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// recursively sort both sublists
	left, right = mergesort(left), mergesort(right)

	return merge(left, right)
}

func merge(a, b []int) []int {
	result := []int{}

	for len(a) > 0 && len(b) > 0 {
		if a[0] <= b[0] {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}

	// consume the remaining elements of the longer list
	for len(a) > 0 {
		result = append(result, a[0])
		a = a[1:]
	}
	for len(b) > 0 {
		result = append(result, b[0])
		b = b[1:]
	}
	return result
}

func main() {
	nums := []int{45, 25, 48, -96, 1, 3, -3, -69, 589,
		479, 56, 54, 852, -258, 88, -395, 42, 89, 313, -8,
		-42, 16, 39, 100, 4, 2, 88, 13, -296, 14, 869, 11}
	nums = mergesort(nums)
	fmt.Println(nums)
}
