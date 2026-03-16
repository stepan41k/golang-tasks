package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(s []int) {
	s[0] = 10
	s = append(s, 4)
	s[1] = 20
}