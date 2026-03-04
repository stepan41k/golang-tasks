package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Слить N каналов в один

func joinChannels(ctx context.Context, chs ...chan int) <-chan int {
	resCh := make(chan int)
	wg := sync.WaitGroup{}

	merger := func(ch <-chan int) {
		select {
		case <-ctx.Done():
			return
		case val, _ := <-ch:
			resCh <- val
		}
	}
		
	for _, v := range chs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			merger(v)
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()
		
	return resCh
}

func main() {
	chs := make([]chan int, 20)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	for i := range chs {
		chs[i] = make(chan int)
	}

	go func() {
		for i, v := range chs {
			v <- i
			close(v)
		}	
	}()
	

	for num := range joinChannels(ctx, chs...) {
       fmt.Println(num)
	}
}