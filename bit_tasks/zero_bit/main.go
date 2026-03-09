package main

import (
	"fmt"
)

func zeroBit(x int, n int) int {
								// (9, 3) 00001001 &  11110111
	return x &^ (1 << n)		// (8, 3) 00001000 &  11110111
	
}

func main() {
	fmt.Println(zeroBit(9, 3))	// 00001001
}