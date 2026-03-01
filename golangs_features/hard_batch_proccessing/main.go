package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type response struct {
	data string
	err  error
}

type request struct {
	id       int
	respChan chan response
}

type BatchLoader struct {
	inputCh chan request
	maxBatchSize int
	linger time.Duration
}


func NewBatchLoader(maxBatchSize int, linger time.Duration) *BatchLoader {
    bl := &BatchLoader{
		inputCh: make(chan request, maxBatchSize),
		maxBatchSize: maxBatchSize,
		linger: linger,
	}

	go func() {
		bl.start()
	}()

	return bl
}


func (l *BatchLoader) Load(ctx context.Context, id int) (string, error) {
	respCh := make(chan response, 1)
	l.inputCh <- request{id: id, respChan: respCh}
	
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case resp := <-respCh:
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			default:
				if resp.err != nil {
					return "", resp.err
				}
				return resp.data, nil
			}
		}
	}
}


func fetchBatch(ids []int) (map[int]string, error) {
	fmt.Printf("--- Fetching batch of size %d: %v ---\n", len(ids), ids)
	results := make(map[int]string)
	for _, id := range ids {
		if id == 3 {
			return nil, errors.New("db connection error")
		}
		results[id] = fmt.Sprintf("Data for %d", id)
	}

	return results, nil
}


func (l *BatchLoader) start() {
	ticker := time.NewTicker(l.linger)
	defer ticker.Stop()
	requestsBatch := make([]request, 0, l.maxBatchSize)

	flush := func() {
		if len(requestsBatch) > 0 {
			cp := make([]int, 0, len(requestsBatch))
			for _, v := range requestsBatch {
				cp = append(cp, v.id)
			}

			results, err := fetchBatch(cp) 

			for _, req := range requestsBatch {
				if err != nil {
					req.respChan <- response{data: "", err: err}
					continue
				}

				val, ok := results[req.id]
				if !ok {
					req.respChan <- response{data: "", err: fmt.Errorf("not found")}
					continue
				}

				req.respChan <- response{data: val, err: nil}
			}

			requestsBatch = requestsBatch[:0]
		}
	}

	for {
		select {
		case <-ticker.C:
			flush()
		case val := <-l.inputCh:
			requestsBatch = append(requestsBatch, val)
			if len(requestsBatch) >= l.maxBatchSize {
				flush()
				ticker.Reset(l.linger)
			}
		}
	}
}


func main() {
	loader := NewBatchLoader(5, 50*time.Millisecond)

	wg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			res, err := loader.Load(context.Background(), id)
			if err != nil {
				fmt.Printf("User %d: Error: %v\n", id, err)
			} else {
				fmt.Printf("User %d: Result: %s\n", id, res)
			}
		}(i)
		time.Sleep(5 * time.Millisecond)
	}
	wg.Wait()
}

