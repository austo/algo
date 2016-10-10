package main

import (
	"fmt"
	"time"
)

func longestPal(s string) string {
	n := len(s)
	cache := make([][]bool, n)

	max := 1
	start := 0

	for i := 0; i < n; i++ {
		cache[i] = make([]bool, n-i)
		cache[i][0] = true
		if i < n-1 && s[i] == s[i+1] {
			cache[i][1] = true
			max = 2
			start = i
		}
	}

	offset := 2

	for offset < n {
		for i := 0; i < n-offset; i++ {
			if cache[i+1][offset-2] && s[i] == s[i+offset] {
				cache[i][offset] = true
				max = offset + 1
				start = i
			}
		}
		offset++
	}
	return s[start : start+max]
}

func longestPalSquare(s string) string {
	n := len(s)
	// var cache [n][n]bool // wish you could do this
	cache := make([][]bool, n)

	max := 1
	start := 0

	for i := 0; i < n; i++ {
		cache[i] = make([]bool, n)
		cache[i][i] = true
		if i < n-1 && s[i] == s[i+1] {
			cache[i][i+1] = true
			max = 2
			start = i
		}
	}

	offset := 2

	for offset < n {
		for i := 0; i < n-offset; i++ {
			if cache[i+1][i+offset-1] && s[i] == s[i+offset] {
				cache[i][i+offset] = true
				max = offset + 1
				start = i
			}
		}
		offset++
	}
	return s[start : start+max]
}

func longestPalFixed(s string) string {
	n := len(s)
	var cache [100][100]bool

	max := 1
	start := 0

	for i := 0; i < n; i++ {
		cache[i][i] = true
		if i < n-1 && s[i] == s[i+1] {
			cache[i][i+1] = true
			max = 2
			start = i
		}
	}

	offset := 2

	for offset < n {
		for i := 0; i < n-offset; i++ {
			if cache[i+1][i+offset-1] && s[i] == s[i+offset] {
				cache[i][i+offset] = true
				max = offset + 1
				start = i
			}
		}
		offset++
	}
	return s[start : start+max]
}

func timeIt(desc string, f func()) {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		f()
	}
	end := time.Now()
	fmt.Printf("%s - %s\n", desc, end.Sub(start))
}

const str = "is this a palindrome: amanaplanacanalpanama?"

func main() {
	timeIt("staggered slice", func() {
		longestPal(str)
	})

	timeIt("square slice", func() {
		longestPalSquare(str)
	})

	timeIt("constant array", func() {
		longestPalFixed(str)
	})
}
