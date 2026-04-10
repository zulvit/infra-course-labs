package main

import (
	"fmt"
	"math"
)

// epsilon - это маленькое число,
// которое определяет "погрешность" сравнения
const epsilon = 1e-9

func floatEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func main() {
	// Классическая ошибка IEEE 754:
	a := 0.1
	b := 0.2
	c := a + b
	d := 0.3
	fmt.Println(c == d)      // false!
	fmt.Printf("%.20f\n", c) // 0.30000000000000004441
	fmt.Printf("%.20f\n", d) // 0.29999999999999998890

	// 0.1 и 0.2 не представимы точно в двоичной системе.

	fmt.Println(floatEqual(0.1+0.2, 0.3, epsilon)) // true

	// Для финансовых расчётов используйте целые числа (суммы в копейках/центах)
	// или библиотеку github.com/shopspring/decimal

	// Бесконечность
	posInf := math.Inf(1)  // +∞
	negInf := math.Inf(-1) // -∞

	fmt.Println(posInf) // true
	fmt.Println(negInf) // true

	var zero float64 = 0.0
	x := 1.0 / zero               // +Inf в рантайме, без паники (!)
	fmt.Println(math.IsInf(x, 1)) // true

	// NaN - Not a Number
	nan := math.NaN()
	fmt.Println(nan == nan)      // false! - стандарт IEEE 754
	fmt.Println(math.IsNaN(nan)) // true - правильный способ

	// Как возникают:
	result := math.Sqrt(-1) // NaN
	fmt.Println(result)
}
