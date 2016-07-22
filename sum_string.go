package main

import "fmt"

const (
	Zero = 48
	Nine = 57
)

func sumString(s string) int {
	sum, i := 0, 0
	for i < len(s) {
		found, moved := consume(s[i:])
		sum += found
		i += moved
	}
	return sum
}

func consume(s string) (int, int) {
	neg := false
	i := 0
	sum := 0
	if s[i] == '-' {
		neg = true
		i++
	}
	for {
		digit, val := isDigit(s[i])
		i++
		if digit {
			sum = sum*10 + val
		} else {
			break
		}
	}
	if neg {
		sum = -sum
	}
	return sum, i
}

func isDigit(b byte) (bool, int) {
	if b >= Zero && b <= Nine {
		return true, int(b - Zero)
	}
	return false, 0
}

func main() {
	fmt.Println(sumString("this is 1 hairy dog and 2 cats"))
	fmt.Println(sumString("this is 11 hairy dogs and 23 cats"))
	fmt.Println(sumString("what is -2 + 3?"))
	fmt.Println(sumString("1 cat and 10 chickens"))
	fmt.Println(sumString("what-you bought 12 pairs of shoes and 3 handbags?"))
}
