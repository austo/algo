package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	r   *rand.Rand
	seq []int
}

// Don't protect the original array
func Constructor(nums []int) Solution {
	return Solution{rand.New(rand.NewSource(time.Now().UnixNano())), nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.seq
}

/** Returns a random shuffling of the array. */
// func (this *Solution) Shuffle() []int {
// 	n := len(this.seq)
// 	r := make([]int, n)

// 	r[0] = this.seq[0]

// 	for i := 1; i < n; i++ {
// 		j := this.r.Intn(i)
// 		fmt.Println(i, j)
// 		r[i] = r[j]
// 		r[j] = this.seq[i]
// 		fmt.Printf("r: %v\n", r)
// 	}

// 	return r
// }

func (this *Solution) Shuffle() []int {
	n := len(this.seq)
	seq := make([]int, n)
	copy(seq, this.seq)

	for i := 1; i < n; i++ {
		j := this.r.Intn(i + 1)
		seq[i], seq[j] = seq[j], seq[i]
	}

	return seq
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	ints := []int{1, 2, 3}
	s := Constructor(ints)
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
}
