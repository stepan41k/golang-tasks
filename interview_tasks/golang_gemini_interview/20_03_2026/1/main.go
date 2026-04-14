package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for i := 0; i < 3; i++ {
		val, ok := <-ch
		fmt.Printf("val: %d, ok: %v\n", val, ok)
	}
}