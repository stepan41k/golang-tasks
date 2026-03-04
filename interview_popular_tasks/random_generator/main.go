package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	resCh := make(chan int, n)

	go func() {
		for i := 0; i < n; i++ {
			resCh <- r.Intn(n)
		}

		close(resCh)
	}()

	return resCh
}

func main() {
	for num := range randomGenerator(10) {
		fmt.Println(num)
	}
}
