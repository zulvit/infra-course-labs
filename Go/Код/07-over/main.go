package main

import (
	"fmt"
	"math"
)

// Константы для границ типов из пакета math
const (
	MaxInt8  = math.MaxInt8  // 127
	MinInt8  = math.MinInt8  // -128
	MaxUint8 = math.MaxUint8 // 255
	MaxInt   = math.MaxInt   // 9223372036854775807 на 64-bit
)

func main() {
	// Go НЕ паникует при переполнении,
	// а просто обрезает результат по модулю диапазона типа.
	var x uint8 = 255
	x++
	fmt.Println(x) // 0 - переполнение!

	var y int8 = 127
	y++
	fmt.Println(y) // -128 - переполнение!

	// Реальная проблема - переполнение в вычислениях
	var a uint8 = 200
	var b uint8 = 100
	sum := a + b     // ожидаем 300, получаем 44 (300 % 256)
	fmt.Println(sum) // 44

	// Проверка перед операцией:
	if int(a)+int(b) > math.MaxUint8 {
		// переполнение - обработать явно
	}
}
