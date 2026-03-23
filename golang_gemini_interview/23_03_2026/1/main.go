package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	go func() {
		fmt.Println("Goroutine started")
		panic("something went wrong")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main finished")
}