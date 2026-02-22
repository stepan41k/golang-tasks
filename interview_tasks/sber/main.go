package main

/*
   1. Сделать структуры Base и Child.
   2. Структура Base должна содержать строковое поле name.
   3. Структура Child должна содержать строковое поле lastName.
   4. Сделать функцию Say у структуры Base, которая распечатывает
        на экране: Hello, %name!
   5. Пронаследовать Child от Base.
   6. Инициализировать экземпляр b1 Base.
        присвоить name значение Parent
   7. Инициализировать экземпляр c1 Сhild.
        присвоить name значение Child
        присвоить lastName значение Inherited
   8. Вызвать у обоих экземпляров метод Say.
   9. Переопределить метод Say для структуры Child, чтобы он выводил
        на экран: Hello, %lastName %name!
   10. Сделать массив, содержащий b1 и c1.
   11. Вызвать Say у всех элементов массива из шага 10.
   12. Сделать метод NewObject для создания экземпляров Base и Child
        в зависимости от входного параметра.
   13. Написать юнит тесты для метода NewObject.
   14. Сделать генератор объектов Base и Child такой, чтобы:
       объекты Base создавались в фоновом потоке с задержкой 1 секунда;
       объекты Child создавались в фоновом потоке с задержкой 2 секунды;
       общее время генерации объектов не превышало 11 секунд;
   15. Сделать асинхронный обработчик сгенерированных объектов такой, чтобы:
       метод Say вызывался в порядке генерации объектов;
       не приводил к утечкам памяти;
*/

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Sayer interface {
	Say()
} 

type Base struct {
	name string
}

func (b Base) Say() {
	fmt.Printf("Hello, %s\n", b.name)
}

type Child struct {
	Base
	lastName string
}

func (c Child) Say() {
	fmt.Printf("Hello %s %s!\n", c.lastName, c.Base.name)
}

func NewObject(str string) Sayer {
	switch str {
	case "Base":
		return Base{name: "Parent"}
	case "Child":
		return Child{lastName: "Inherited", Base: Base{name: "Child"}}
	default:
		return nil
	}
}

func NewObjectGenerator(ctx context.Context) (<-chan Sayer) {
	resChan := make(chan Sayer, 1)

	wg := &sync.WaitGroup{}

	tickerBase := time.NewTicker(1 * time.Second)
	tickerChild := time.NewTicker(2 * time.Second)

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			tickerBase.Stop()
		}()

		for {
			select {
			case <- ctx.Done():
				return
			case <-tickerBase.C:
				resChan <- Base{name: "Hello Base from generator"}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			tickerChild.Stop()
		}()

		for {
			select {
			case <- ctx.Done():
				return
			case <-tickerChild.C:
				resChan <- Child{Base: Base{name: "Hello Inherited from generator"}, lastName: "Hello Child from generator"}
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resChan)
	}()
	
	return resChan
}

func AsyncHandler(ch <-chan Sayer) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for e := range ch {
			e.Say()
		}
	}()

	wg.Wait()
}

func main() {
	// b1 := &Base{name: "Parent"}
	// c1 := &Child{Base: Base{name: "Child"}, lastName: "Inherited"}

	// b1.Say()
	// c1.Say()

	// relArr := []Sayer{b1, c1}
	// for _, v := range relArr {
	// 	v.Say()
	// }


	ctx, cancel := context.WithTimeout(context.Background(), 11 * time.Second)
	defer cancel()

	sayers := NewObjectGenerator(ctx)

	AsyncHandler(sayers)
}	

//TODO: Odner
// package main

// import (
// 	"context"
// 	"time"
// 	"fmt"
// 	"sync"
// )

// type Sayer interface {
// 	Say()
// }

// type Base struct {
// 	name string
// }

// func (b Base) Say() {
// 	fmt.Printf("Hello, %s\n", b.name)
// }

// type Child struct {
// 	Base
// 	lastName string
// }

// func (c Child) Say() {
// 	fmt.Printf("Hello %s %s!\n", c.lastName, c.Base.name)
// }

// func NewObject(str string) Sayer {
// 	switch str {
// 	case "Base":
// 		return Base{name: "Parent"}
// 	case "Child":
// 		return Child{lastName: "Inherited", Base: Base{name: "Child"}}
// 	default:
// 		return nil
// 	}
// }

// func Generator(ctx context.Context) <-chan Sayer {
// 	ch := make(chan Sayer)

// 	worker := func(ticker *time.Ticker, typeObject string) {
// 		defer ticker.Stop()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case <-ticker.C:
// 				ch <- NewObject(typeObject)
// 			}
// 		}
// 	}

// 	wg := sync.WaitGroup{}

// 	go func() {
// 		defer func() {
// 			wg.Wait()
// 			close(ch)
// 		}()
		
// 		wg.Go(func() {worker(time.NewTicker(1*time.Second), "Base")})
		
// 		wg.Go(func() {worker(time.NewTicker(2*time.Second), "Child")})
// 	}()

// 	return ch
// }


// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
// 	defer cancel()

// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
	
// 	go func() {
// 		defer wg.Done()

// 		for newObj := range Generator(ctx) {
// 			newObj.Say()
// 		}
// 	}()

// 	wg.Wait()
// }
