package main

import (
	"fmt"
	"sync/atomic"
)

// Сделать конвейер чисел
// Даны два канала. В первый пишутся числа. Нужно, чтобы числа читались из первого по мере поступления, что-то с ними происходило (допустим, возводились в квадрат) и результат записывался во второй канал.

var (
	lenCh = 10
)

func main() {
	inputCh := make(chan int, lenCh)
	outputCh := make(chan int, lenCh)
	var counter int64 = 0

	go func() {
		for i := 0; i < 100; i++ {
			inputCh <- i
		}

		close(inputCh)
	}()

	go func() {
		for v := range inputCh {
			atomic.AddInt64(&counter, 1)
			outputCh <- v * v
		}

		close(outputCh)
	}()

	for v := range outputCh {
		fmt.Println(v)
	}

	fmt.Println(counter)
}
