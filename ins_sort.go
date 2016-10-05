package main

import "fmt"

func inssort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

var coll []int = []int{
	12, 15, 7, 3, 86, 18, 100, 302, 98, 40,
}

func main() {
	inssort(coll)
	fmt.Println(coll)
}
