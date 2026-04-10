package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Объявление
	var c1 complex64 = 3 + 4i
	var c2 complex128 = 3.14 + 2.72i
	c3 := 5 + 3i            // тип complex128 по умолчанию
	c4 := complex(2.0, 3.0) // через встроенную функцию

	fmt.Printf("c1: %v, c2: %v, c3: %v, c4: %v\n", c1, c2, c3, c4)

	// Извлечение частей
	fmt.Println(real(c3)) // 5
	fmt.Println(imag(c3)) // 3

	// Арифметика
	sum := c1 + complex64(3+2i)
	fmt.Println(sum) // (6+6i)

	// Пакет math/cmplx
	fmt.Println(cmplx.Abs(c3))  // 5.83... = sqrt(25+9)
	fmt.Println(cmplx.Sqrt(-1)) // (0+1i) - мнимая единица
}
