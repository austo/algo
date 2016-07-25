package main

import (
	"flag"
	"fmt"
)

var n = flag.Int("n", 10, "desired sum")

func main() {
	flag.Parse()
	handle(*n)
}

func handle(sum int) {
	var next int
	m := map[int]int{}

	for {
		_, err := fmt.Scanf("%d", &next)
		if err != nil {
			break
		}
		compliment := sum - next
		if v, ok := m[compliment]; ok {
			fmt.Printf("%d + %d = %d\n", compliment, next, sum)
			if v == 1 {
				delete(m, compliment)
			} else {
				m[compliment]--
			}
		} else {
			m[next]++
		}
	}
}
