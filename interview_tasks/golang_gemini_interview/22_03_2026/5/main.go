package main

import (
	"context"
	"fmt"
	"time"
)

// Функция читает числа из канала in и группирует их в слайсы (батчи) размером size.
// Как только батч заполнен, он отправляется в возвращаемый канал.
// Если входной канал in закрывается, все накопленные на этот момент элементы должны быть отправлены в выходной канал одним (возможно, неполным) батчем.
// После отправки последнего батча выходной канал должен закрыться.
// Решение должно быть эффективным и не допускать утечек горутин.

func Batch(in <-chan int, size int) <-chan []int {
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	batch := make([]int, 0, size)
	resCh := make(chan []int)

	flush := func() {
		if len(batch) > 0 {
			cpBatch := make([]int, len(batch))
			copy(cpBatch, batch)
			resCh <- cpBatch
			batch = batch[:0]
		}
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				flush()
				close(resCh)
				return
			case val, ok := <-in:
				if !ok {
					flush()
					close(resCh)
					return
				}

				batch = append(batch, val)
				if len(batch) >= size {
					flush()
				}
			}
		}
	}()

	return resCh
}

func main() {
	ch := make(chan int, 1)
	val := 0

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ticker.C:
				ch <- val
				val++
			case <-ctx.Done():
				close(ch)
				return
			}
		}
	}()

	for batch := range Batch(ch, 6) {
		fmt.Println(batch)
	}
}
