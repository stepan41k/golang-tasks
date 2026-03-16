package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	countWorkers = 3
	countTasks   = 29
)


func ExecuteTasks(tasks []func(), n int, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	taskCh := make(chan func())
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case t, ok := <-taskCh:
					if !ok {
						return
					}
					t()
					fmt.Printf("hello from %d\n", n)
				}
			}

		}(i)
	}

	go func() {
		for _, t := range tasks {
			select {
			case <-ctx.Done():
				break
			case taskCh <- t: 
			}
		}

		close(taskCh)
	}()

	wg.Wait()

	return ctx.Err()
}

func main() {
	task := func() {}

	tasks := make([]func(), countTasks)
	executeTimeout := time.Second
	wg := sync.WaitGroup{}

	for i := 0; i < countTasks; i++ {
		tasks[i] = task
	}

	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ExecuteTasks(tasks, n, executeTimeout)
		}(i)
	}

	wg.Wait()

}
