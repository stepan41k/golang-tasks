package main

import (
	"context"
	"fmt"
	"sync/atomic"

	"math/rand"
	"sync"

	"time"
)

type Res struct {
	ID int32
}

type Pool struct {
	resources chan Res
	semaphore chan struct{}
	factory func() Res
}


func NewPool(maxSize int, factory func() Res) *Pool {	
    newPool := &Pool{
		resources: make(chan Res, maxSize),
		semaphore: make(chan struct{}, maxSize),
		factory: factory,
	}

	return newPool
}


func (p *Pool) Acquire(ctx context.Context) (Res, error) {

	select {
	case val := <-p.resources:
		return val, nil
	default:
	}

	select {
	case <-ctx.Done():
		return Res{}, ctx.Err()
	case p.semaphore <- struct{}{}:
		return p.factory(), nil
	default:
		select {
		case val := <-p.resources:
			return val, nil
		case <-ctx.Done():
			return Res{}, ctx.Err()
		}
	}
}

func (p *Pool) Release(res Res) {
	p.resources <- res
}

func main() {
	var counter int32 = 0
	wg := sync.WaitGroup{}

    p := NewPool(10, func() Res {
		atomic.AddInt32(&counter, 1)
        return Res{ID: counter}
    })

    ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := p.Acquire(ctx)
			fmt.Println("Acquire:", res)
    		if err == nil {
				time.Sleep(time.Duration(rand.Intn(4) * int(time.Second)))
				fmt.Println("Release:", res)
        		p.Release(res)
    		}
		}()
	}

	wg.Wait()
}