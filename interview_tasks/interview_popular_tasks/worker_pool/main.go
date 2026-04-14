package main

import (
	"fmt"
	"sync"
)

// Написать WorkerPool с заданной функцией
// Нам нужно разбить процессы на несколько горутин — при этом не создавать новую горутину каждый раз, а просто переиспользовать уже имеющиеся.

func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
        results <- f(j)
    }
}

func main() {
	const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    multiplier := func(x int) int {
	    return x * 10
    }

	wg := sync.WaitGroup{}

	wg.Add(3)
    for i := 1; i <= 3; i++ {
        go func() {
			defer wg.Done()
			worker(i,  multiplier, jobs, results)
		}()
    }

	go func() {
		for j := 1; j <= numJobs; j++ {
        	jobs <- j
    	}

		close(jobs)
	}()
    
	go func() {
		wg.Wait()
		close(results)
	}()

    for v := range results {
		fmt.Println(v)
	}
}

