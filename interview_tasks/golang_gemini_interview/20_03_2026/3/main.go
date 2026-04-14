package main

import (
	"fmt"
	"sync"
)

// Тебе нужно написать функцию Merge, которая принимает на вход слайс каналов []<-chan int и возвращает один канал <-chan int.

// Все значения из входных каналов должны быть пересланы в результирующий канал.
// Результирующий канал должен закрыться только тогда, когда закроются все входные каналы.
// Порядок значений не важен.
// Постарайся написать решение, которое не будет «плодить» лишние горутины (на каждый входной канал — одна горутина — это нормально).

func Merge(channels ...<-chan int) <-chan int {
	resCh := make(chan int, len(channels))
	wg := sync.WaitGroup{}

	merger := func(ch <-chan int) {
		for v := range ch {
			resCh <- v
		}
	}

    for _, v := range channels {
		wg.Add(1)
		go func(v <-chan int) {
			defer wg.Done()
			merger(v)
		}(v)
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	return resCh
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	close(ch1)
	close(ch2)
	close(ch3)

	for v := range Merge(ch1, ch2, ch3) {
		fmt.Println(v)
	}
}