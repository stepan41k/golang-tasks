package main

import "fmt"

func main() {
	var ch chan int 

	select {
	case v := <-ch:
		fmt.Println("Received:", v)
	case ch <- 1:
		fmt.Println("Sent")
	default:
		fmt.Println("Default")
	}
}