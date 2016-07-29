package main

import "fmt"

var (
	ctr  int
	dict map[string]bool = map[string]bool{
		"anti":                  true,
		"dis":                   true,
		"disestablish":          true,
		"establish":             true,
		"establishment":         true,
		"establishmentarian":    true,
		"establishmentarianism": true,
		"ment":                  true,
		"ism":                   true,
	}
)

func main() {
	m := substrings("antidisestablishmentarianism")
	for k, _ := range m {
		if dict[k] {
			fmt.Println(k)
		}
	}
	fmt.Printf("%d operations\n", ctr)
}

func substrings(s string) map[string]struct{} {
	m := map[string]struct{}{}
	substringsR(s, m)
	return m
}

func substringsR(s string, m map[string]struct{}) {
	if _, ok := m[s]; ok {
		return
	}
	ctr++
	m[s] = struct{}{}
	if len(s) > 1 {
		substringsR(s[1:], m)
		substringsR(s[0:len(s)-1], m)
	}
}
