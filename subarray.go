package main

import "fmt"

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	pretty(subarrays(arr, 2, 8))
}

func subarrays(a []int, low, high int) [][2]int {
	result := [][2]int{}
	currSum, start, n := a[0], 0, len(a)

	for i := 1; i <= n; i++ {

		for (currSum > high || currSum < low) && start < i-1 {
			currSum -= a[start]
			start++
		}

		if currSum > low && currSum < high {
			result = append(result, [2]int{start, i - 1})
		}

		if i < n {
			currSum += a[i]
		}
	}

	currSum -= a[start]
	start++

	for currSum > low && currSum < high && start < n {
		result = append(result, [2]int{start, n - 1})
		currSum -= a[start]
		start++
	}
	return result
}

func pretty(vv [][2]int) {
	fmt.Print("{")

	for i, v := range vv {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("(%d,%d)", v[0], v[1])
	}

	fmt.Println("}")
}
