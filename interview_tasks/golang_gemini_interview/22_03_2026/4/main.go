package main

import (
	"context"
	"time"
	"fmt"
)

// // Нужно запустить все поисковики одновременно.
// // Как только первый из них вернул результат, функция должна немедленно вернуть этот результат.
// // Важно: Остальные поисковики должны получить сигнал об отмене (представим, что они принимают context.Context), чтобы не занимать ресурсы.
// // Утечек горутин быть не должно.

const (
	getResponseTimeount = 6 * time.Second
)

func GetFirstResponse(ctx context.Context, searchers []func(ctx context.Context) string) (res string) {
	ctx, cancel := context.WithTimeout(ctx, getResponseTimeount)
	defer cancel()

	resCh := make(chan string, len(searchers))

	for _, v := range searchers {
		go func() {
			select {
			case <-ctx.Done():
				return
			default:
				val := v(ctx)
				select {
				case <-ctx.Done():
					return
				case resCh <- val:
				}
			}
		}()
	}

	select {
	case <-ctx.Done():
		return "context done"
	case res = <-resCh:
		cancel()
	}

	return res
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6 * time.Second)
	defer cancel()

	f1 := func(ctx context.Context) string {
		time.Sleep(220 * time.Millisecond)
		return "Hi there from f1"
	}

	f2 := func(ctx context.Context) string {
		time.Sleep(220 * time.Millisecond)
		return "Hi there from f2"
	}

	f3 := func(ctx context.Context) string {
		time.Sleep(220 * time.Millisecond)
		return "Hi there from f3"
	}

	f4 := func(ctx context.Context) string {
		time.Sleep(220 * time.Millisecond)
		return "Hi there from f4"
	}

	f5 := func(ctx context.Context) string {
		time.Sleep(220 * time.Millisecond)
		return "Hi there from f5"
	}

	arg := []func(ctx context.Context) string{f1, f2, f3, f4, f5}

	fmt.Println(GetFirstResponse(ctx, arg))
}
