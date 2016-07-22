package main

import "fmt"

const (
	MaxInt32      = 1<<31 - 1
	MinInt32      = -1 << 31
	MaxDecimalLen = 10
	Zero          = 48
	Minus         = 45
)

var Numerals = [...]byte{
	0: 48,
	1: 49,
	2: 50,
	3: 51,
	4: 52,
	5: 53,
	6: 54,
	7: 55,
	8: 56,
	9: 57,
}

func main() {
	ints := []int{-123, 456, 42, 0}
	for _, v := range ints {
		fmt.Printf("%d becomes %q\n", v, intToStr(v))
	}
}

func intToStr(k int) string {
	if k == 0 {
		return "0"
	}
	neg := false
	digits := make([]byte, 0, MaxDecimalLen+1)
	if k < 0 {
		k = -k
		neg = true
	}
	for k != 0 {
		r := k % 10
		digits = append(Numerals[r:r+1], digits...)
		// digits = append([]byte{byte((k % 10) + Zero)}, digits...)
		k /= 10
	}
	if neg {
		digits = append([]byte{Minus}, digits...)
	}
	return string(digits)
}
