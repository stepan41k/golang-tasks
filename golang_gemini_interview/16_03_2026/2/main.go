package main

import (
	"fmt"
	"sync"
)

func main() {
	var m map[string]int
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m[fmt.Sprintf("key-%d", n)] = n
		}(i)
	}

	wg.Wait()
	fmt.Println("Done", len(m))
}