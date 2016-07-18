/*
	There are N coins with coordinates (x, y) where x > 0 and y > 0
	You start at (0, 0) and you can only do steps of form (dx, dy) where dx > 0 and dy > 0
	Print the maximum number of coins that you can collect.

	Clarification: you can do as many moves as you wish, the point is to collect maximum number of coins. If you are located at position (a, b) you may jump to position (a+dx, b+dy) for all dx > 0 and dy > 0

	@krbchd: Your algorithm may output incorrect values. Suppose there are points (5, 7), (5, 8), (5, 9) for y coordinates LIS will output 7, 8, 9, however since these points are on the same x axis, you can choose only one of them.

	0 0 1 1 [[3,1], [1,1], [2,2], [1,2], [3,3], [2,3]]
	0 1 1 0
	0 1 0 1
	0 0 0 0

	Brute force:
	1. create n * n matrix with number of coins stored at each coordinate
	2. starting at given coordinate, find the coordinate giving max coins via recursion


	Thoughts:
	1. sort coins by ascending y coordinate and descending x coordinate
			[[3,1], [1,1], [2,2], [1,2], [3,3], [2,3]]



*/

package main

import "fmt"

// func buildCoinMatrix(points [][2]int) [][]int {
// 	max := 0
// 	for _, point := range points {
// 		for _, n := range point {
// 			if n > max {
// 				max = n
// 			}
// 		}
// 	}
// 	m := buildMatrix(max+1, 0)
// 	for _, point := range points {
// 		m[point[1]][point[0]]++
// 	}
// 	fmt.Println(m)
// 	return m
// }

// func buildMatrix(n int, pop int) [][]int {
// 	m := make([][]int, n)
// 	rows := make([]int, n*n)
// 	if pop != 0 {
// 		for i := 0; i < n*n; i++ {
// 			rows[i] = pop
// 		}
// 	}
// 	for i := range m {
// 		m[i], rows = rows[:n], rows[n:]
// 	}
// 	return m
// }

func buildMap(points [][2]int) (map[string]int, int) {
	m := map[string]int{}
	max := 0
	for _, point := range points {
		for _, v := range point {
			if v > max {
				max = v
			}
		}
		cacheKey := fmt.Sprintf("%d-%d", point[1], point[0])
		m[cacheKey] = m[cacheKey] + 1
	}
	return m, max + 1
}

func maxCoins(m map[string]int, row, col, n int, cache map[string]int) int {
	if row >= n || col >= n {
		panic("invalid index")
	}
	cacheKey := fmt.Sprintf("%d-%d", row, col)
	if v, ok := cache[cacheKey]; ok {
		return v
	}

	max := 0
	for dx := 1; col+dx < n; dx++ {
		for dy := 1; row+dy < n; dy++ {
			newMax := maxCoins(m, row+dy, col+dx, n, cache)
			if newMax > max {
				max = newMax
			}
		}
	}
	max += m[cacheKey]
	cache[cacheKey] = max
	return max
}

func main() {
	points := [][2]int{
		[2]int{0, 0},
		[2]int{1, 1},
		[2]int{1, 2},
		[2]int{2, 2},
		[2]int{2, 3},
		[2]int{3, 1},
		[2]int{3, 3},
	}
	m, n := buildMap(points)
	fmt.Println(maxCoins(m, 0, 0, n, map[string]int{}))
}
