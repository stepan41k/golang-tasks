package main

import (
	"fmt"
)

// Реализуй функцию isPowerOfTwo(n int) bool.
// Функция должна возвращать true, если число n является степенью двойки (1, 2, 4, 8, 16...), и false в противном случае.
// Условие:
// Нельзя использовать циклы (for).
// Нельзя использовать библиотеку math.
// Нужно использовать одну хитрую битовую операцию.
// Считаем, что n > 0.
// Пример:
// Вход: n = 8 (1000₂) → Результат: true
// Вход: n = 7 (0111₂) → Результат: false

// Подсказка: 8 = 1000   7 = 0111

func isPowerOfTwo(n int) bool {
	if n & (n-1) == 0 { 
		return true
	}  	
	return false
}					

func main() {
	fmt.Println(isPowerOfTwo(10))
	fmt.Println(isPowerOfTwo(8))
	fmt.Println(isPowerOfTwo(15))
	fmt.Println(isPowerOfTwo(16))
}