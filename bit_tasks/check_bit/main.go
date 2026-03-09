package main

import (
	"fmt"
)

// Реализуй функцию 
// Функция должна вернуть true, если i-ый бит числа n равен 1, и false, если он равен 0.
// Пример:
// Вход: n = 10 (1010₂), i = 1 → Результат: true
// Вход: n = 10 (1010₂), i = 2 → Результат: false

func hasBit(n int, i int) bool {
	if n & (1 << i) != 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(hasBit(10, 1))
	fmt.Println(hasBit(10, 2))
	fmt.Println(hasBit(12, 3))
}