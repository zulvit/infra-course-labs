package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// bool — 1 байт в памяти
	var isActive bool // false — нулевое значение
	isReady := true
	fmt.Println(unsafe.Sizeof(isActive)) // 1
	fmt.Println(isActive, isReady)       // false true
	// Логические операторы:
	a, b := true, false

	fmt.Println(a && b) // false — AND: оба должны быть true
	fmt.Println(a || b) // true  — OR: хотя бы один true
	fmt.Println(!a)     // false — NOT: инверсия

	// Таблица истинности AND (&&):
	// true  && true  = true
	// true  && false = false
	// false && true  = false
	// false && false = false

	// Таблица истинности OR (||):
	// true  || true  = true
	// true  || false = true
	// false || true  = true
	// false || false = false
}
