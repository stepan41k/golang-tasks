package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Сделать кастомную waitGroup на семафоре
// Семафор можно легко получить из канала. Чтоб не аллоцировать лишние данные, будем складывать туда пустые структуры.

type CustomWaitGroup struct {
	size   atomic.Int64
	doneCh    chan struct{}
	mu sync.Mutex
}

func NewCustomWG() *CustomWaitGroup {
	return &CustomWaitGroup{
		doneCh:    make(chan struct{}),
	}
}

func (cwg *CustomWaitGroup) Add(delta int64) {
	cwg.mu.Lock()
	defer cwg.mu.Unlock()

	newSize := cwg.size.Add(delta)

	if newSize < 0 {
		panic("negative waitgroup counter")
	}

	select {
	case <-cwg.doneCh:
		cwg.doneCh = make(chan struct{})
	default:
	}
}

func (cwg *CustomWaitGroup) Done() {
	newSize := cwg.size.Add(-1)
	
	if newSize < 0 {
		panic("negative waitgroup counter")
	}

	if newSize == 0 {
		cwg.mu.Lock()
		close(cwg.doneCh)
		cwg.mu.Unlock()
	}
}

func (cwg *CustomWaitGroup) Wait() {
	cwg.mu.Lock()
	ch := cwg.doneCh
	cwg.mu.Unlock()

	if cwg.size.Load() == 0 {
		return
	}

	<-ch
}

func main() {
	wg := NewCustomWG()

	for i := 0; i < 10; i ++ {
		wg.Add(1)

		go func(d int) {
			defer wg.Done()
			fmt.Printf("Hi there from goroutine №%d\n", d)
		}(i)
	}

	wg.Wait()

	fmt.Println("success")
}
