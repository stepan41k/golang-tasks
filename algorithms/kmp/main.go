package main

import "fmt"

func computePrefix(pattern string) []int {
	pi := make([]int, len(pattern))
	j := 0

	for i := 1; i < len(pattern); i++ {
		if j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}

		pi[i] = j
	}

	return pi
}

func KMP(text, pattern string) []int {
	if len(pattern) == 0 {
		return nil
	}

	res := []int{}
	pi := computePrefix(pattern)
	j := 0

	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != text[j] {
			j = pi[j-1]
		}

		if text[i] == pattern[j] {
			j++
		}

		if j == len(pattern) {
			res = append(res, i - len(pattern) + 1)
			j = pi[j-1]
		}
	}

	return res
}

func main() {
	fmt.Println(KMP("ABACABABEBRAABACABA", "ABACABA"))
}
