package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println(a == b) // false — равно
	fmt.Println(a != b) // true  — не равно
	fmt.Println(a < b)  // true  — меньше
	fmt.Println(a > b)  // false — больше
	fmt.Println(a <= b) // true  — меньше или равно
	fmt.Println(a >= b) // false — больше или равно

	// Что сравнимо через ==:
	// Числа: int, float, etc.
	// Строки: посимвольно (побайтово)
	// bool
	// Указатели: совпадение адреса
	// Структуры (если все поля сравнимы)

	s1, s2 := "hello", "Hello"
	fmt.Println(s1 == s2) // false: 'h' != 'H'

	// Лексикографическое сравнение строк:
	fmt.Println("apple" < "banana") // true — 'a' < 'b'
	fmt.Println("abc" < "abd")      // true — 'c' < 'd'
}
