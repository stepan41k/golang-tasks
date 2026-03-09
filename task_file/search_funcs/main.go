package main

import (
	"context"
	"errors"
	"fmt"
)

// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
// Когда получаем первый успешный результат - отдаем его сразу. Если все SearchFunc отработали
// с ошибкой - отдаем последнюю полученную ошибку

type Result struct{
	data string
}


type SearchFunc func(ctx context.Context, query string) (Result, error)


func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	if len(sfs) == 0 {
		return Result{}, errors.New("no search functions proviede")
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	type response struct {
		res Result
		err error
	}
	
	resCh := make(chan response, len(sfs))

	for _, sf := range sfs {
		go func(f SearchFunc) {
			res, err := f(ctx, query)
			resCh <- response{res, err}
		}(sf)
	}

	var lastError error
	errCount := 0

	for i := 0; i < len(sfs); i++ {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			return Result{}, ctx.Err()
		case r := <-resCh:
			if r.err == nil {
				cancel()
				return r.res, nil
			}

			lastError = r.err
			errCount++
		}
	}

	return Result{}, lastError
}

