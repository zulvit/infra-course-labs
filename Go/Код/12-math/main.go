package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Округление
	math.Ceil(2.3)  // 3.0 - вверх до целого
	math.Floor(2.9) // 2.0 - вниз до целого
	math.Round(2.5) // 3.0 - стандартное округление
	math.Trunc(2.9) // 2.0 - отбросить дробную часть

	// Модуль и знак
	math.Abs(-5.5)           // 5.5
	math.Copysign(3.0, -1.0) // -3.0

	// Степени и корни
	math.Sqrt(16.0) // 4.0
	math.Cbrt(27.0) // 3.0
	math.Pow(2, 10) // 1024.0
	math.Pow10(3)   // 1000.0

	// Логарифмы
	math.Log(math.E) // 1.0 (натуральный)
	math.Log2(8)     // 3.0
	math.Log10(1000) // 3.0

	// Константы
	// math.Pi         // 3.141592653589793
	// math.E          // 2.718281828459045
	// math.MaxFloat64 // 1.7976931348623157e+308
	// math.MaxInt     // 9223372036854775807 (на 64-bit)
	// math.MinInt     // -9223372036854775808

	f := 123456.789

	fmt.Printf("%f\n", f)     // 123456.789000  - стандартный
	fmt.Printf("%.2f\n", f)   // 123456.79      - 2 знака после точки
	fmt.Printf("%e\n", f)     // 1.234568e+05   - научная нотация
	fmt.Printf("%E\n", f)     // 1.234568E+05   - заглавная E
	fmt.Printf("%g\n", f)     // 123456.789     - компактная форма
	fmt.Printf("%9.2f\n", f)  // " 123456.79"   - ширина 9, выравнивание вправо
	fmt.Printf("%-9.2f\n", f) // "123456.79 "   - выравнивание влево

	// Форматирование в строку:
	s := fmt.Sprintf("%.4f", math.Pi) // "3.1416"
	fmt.Println(s)

	// strconv для высокой производительности:
	s2 := strconv.FormatFloat(math.Pi, 'f', 4, 64) // "3.1416"
	// Аргументы: значение, формат ('f','e','g'), точность, разрядность (32/64)
	fmt.Println(s2)

	// Разбор строки в float:
	f2, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(f2) // 3.14
}
