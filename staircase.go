package main

import (
	"flag"
	"os"
)

var (
	c = flag.String("c", "#", "char to use")
	h = flag.Int("h", 10, "height")
)

func main() {
	flag.Parse()
	staircase(*h, *c)
}

func staircase(n int, p string) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= n-i {
				os.Stdout.WriteString(" ")
			} else {
				os.Stdout.WriteString(p)
			}
		}
		os.Stdout.WriteString("\n")
	}
}
