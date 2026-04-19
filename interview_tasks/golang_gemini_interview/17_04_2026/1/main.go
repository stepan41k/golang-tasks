package main

import (
	"fmt"
	"time"
)

// Твое задание:
// Напиши функцию RunTasks, которая принимает:
// tasks []func() error — список функций для выполнения.
// concurrency int — максимальное количество одновременно работающих горутин.
// timeout time.Duration — общий таймаут на выполнение всей пачки задач.
//
// Требования:
// Функция должна возвращать слайс из ошибок []error (результаты каждой задачи). Если задача выполнилась успешно, в слайсе должен быть nil.
// Должно соблюдаться ограничение concurrency.
// Если время timeout вышло, все незапущенные задачи не должны начинаться, а запущенные должны (по возможности) прерваться, и функция должна вернуть управление.
// Важно: Нужно минимизировать аллокации и утечки горутин.

func RunTasks(tasks []func() error, concurrency int, timeout time.Duration) []error {
	type task struct {
		fn         func() error
		taskNumber int
	}

	ch := make(chan task, len(tasks))
	results := []error{}

	for i := 0; i < concurrency; i++ {
		go func() {
			select {
			case val := <-ch:
				if val.taskNumber == 10 {
					results = append(results, fmt.Errorf("10th task"))
				} else {
					results = append(results, nil)
				}
			}
		}()
	}

	for i := 0; i < len(tasks); i++ {
		ch <- task{fn: tasks[i], taskNumber: i}
	}
	
	return results
}


func main() {
	fmt.Println("Hi")
	
	nil := 10
	
	fmt.Println(nil)
}
