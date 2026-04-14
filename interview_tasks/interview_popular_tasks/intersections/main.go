package main

import "fmt"

// На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение

func intersection(a, b []int) []int {
	ints := map[int]int{}
	result := []int{}

	for _, v := range a {
		ints[v]++
	}

	for _, v := range b {
		val, _ := ints[v]
		if val > 0 {
			ints[v]--
			result = append(result, v)
		}
	}

	return result
}


func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
}