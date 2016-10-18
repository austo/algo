package main

/*
 * The goal of this problem is to implement a variant of the 2-SUM algorithm (covered in the Week 6 lecture on hash table applications).
 * The file contains 1 million integers, both positive and negative (there might be some repetitions!).This is your array of integers, with the ith row of the file specifying the ith entry of the array.
 * Your task is to compute the number of target values t in the interval [-10000,10000] (inclusive) such that there are distinct numbers x,y in the input file that satisfy x+y=t. (NOTE: ensuring distinctness requires a one-line addition to the algorithm from lecture.)
 * Write your numeric answer (an integer between 0 and 20001) in the space provided.
 * OPTIONAL CHALLENGE: If this problem is too easy for you, try implementing your own hash table for it. For example, you could compare performance under the chaining and open addressing approaches to resolving collisions.
 */

import (
	"fmt"
	"sort"
)

type doubles []int64

func (d doubles) Len() int           { return len(d) }
func (d doubles) Less(i, j int) bool { return d[i] < d[j] }
func (d doubles) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func main() {
	handle()
}

func handle() {
	input := make([]int64, 0, 1000000)

	var next int64

	for {
		_, err := fmt.Scanf("%d", &next)
		if err != nil {
			break
		}
		input = append(input, next)
	}

	sort.Sort(doubles(input))

	var min int64 = -10000
	var max int64 = 10000

	var left, right int64

	found := map[int64]bool{}

	left = 0
	right = int64(len(input) - 1)

	for left < right {
		t := input[left] + input[right]
		if t < min {
			left++
		} else if t > max {
			right--
		} else {
			found[t] = true
			r := right - 1
			for input[left]+input[r] >= min {
				if input[left] != input[r] {
					found[input[left]+input[r]] = true
				}
				r--
			}
			l := left + 1
			for input[l]+input[right] <= max {
				if input[l] != input[right] {
					found[input[l]+input[right]] = true
				}
				l++
			}
			left++
			right--
		}
	}

	fmt.Println(len(found))
}
