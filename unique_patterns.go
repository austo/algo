package main

import "fmt"

/*
	Given a Pattern and a dictionary, print out all the strings that match the pattern,
	where a character in the pattern is mapped uniquely to a character in the dictionary.

	e.g.
	1. ("abc" , <"cdf", "too", "hgfdt" ,"paa">) -> output = "cdf"
	2. ("acc" , <"cdf", "too", "hgfdt" ,"paa">) -> output = "too", "paa"
*/

func main() {
	m := map[string][]string{
		"abc": []string{"cdf", "too", "hgfdt", "paa"},
		"acc": []string{"cdf", "too", "hgfdt", "paa"},
	}
	for k, _ := range m {
		fmt.Println(uniquePatterns(m, k))
	}
}

func uniquePatterns(m map[string][]string, pat string) []string {
	var result []string
	h := hash(pat)

	for _, s := range m[pat] {
		if hash(s) == h {
			result = append(result, s)
		}
	}

	return result
}

func hash(s string) string {
	var start byte
	var result []byte
	m := map[byte]byte{}
	for _, c := range []byte(s) {
		if v, ok := m[c]; ok {
			result = append(result, v)
		} else {
			m[c] = start
			result = append(result, start)
			start++
		}
	}
	return string(result)
}
