package main

import "fmt"

type elem struct {
	id     int
	weight float32
}

var (
	first []elem = []elem{
		elem{1, 1.0079},
		elem{2, 4.0026},
		elem{3, 6.941},
		elem{4, 9.0122},
		elem{5, 10.811},
		elem{6, 12.011},
		elem{7, 14.007},
		elem{8, 15.999},
	}
	second []elem = []elem{
		elem{6, 12.012},
		elem{7, 14.006},
		elem{8, 16.01},
		elem{9, 18.998},
		elem{10, 20.180},
		elem{11, 22.990},
		elem{12, 24.305},
		elem{13, 26.982},
		elem{14, 28.086},
	}
)

func main() {
	fmt.Println(merge(first, second))
}

// merge
func merge(a, b []elem) []elem {
	result := []elem{}
	m := map[int]int{}

	for len(a) > 0 && len(b) > 0 {
		var e elem
		if a[0].weight < b[0].weight {
			e = a[0]
			a = a[1:]
		} else {
			e = b[0]
			b = b[1:]
		}
		result = insert(result, e, m)
	}

	for len(a) > 0 {
		e := a[0]
		a = a[1:]
		result = insert(result, e, m)
	}
	for len(b) > 0 {
		e := b[0]
		b = b[1:]
		result = insert(result, e, m)
	}

	return result
}

func insert(elems []elem, e elem, m map[int]int) []elem {
	if i, ok := m[e.id]; ok {
		elems[i].weight = (elems[i].weight + e.weight) / 2
	} else {
		m[e.id] = len(elems)
		elems = append(elems, e)
	}
	return elems
}
