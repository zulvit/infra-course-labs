package main

import (
	"fmt"
	"strconv"
)

func main() {
	// В Go НЕТ неявных преобразований типов
	var i int = 42
	var f float64 = float64(i) // явное преобразование
	var u uint = uint(f)

	fmt.Println(i, f, u)

	// При сужении типа - возможна потеря данных (без ошибки!):
	var a int32 = 100
	var c int8 = int8(a) // OK, но...
	fmt.Println(c)       // 100 - всё ещё OK, но...
	big := 200
	fmt.Println(int8(big)) // -56 - тихая потеря!

	// string и []byte / []rune:
	s := "Hello"
	b := []byte(s)  // копия байт
	s2 := string(b) // обратно в строку
	r := []rune(s)  // в руны
	s3 := string(r) // обратно
	fmt.Println(s, b, s2, r, s3)

	// Числа и строки - ТОЛЬКО через strconv:
	n := 42
	// string(n)           // "∗" - руна с кодом 42, НЕ "42"!
	s4 := strconv.Itoa(n) // "42" - правильно
	fmt.Println(s4)

	// float to int
	f2 := 3.14
	i2 := int(f2) // 3 - дробная часть отбрасывается, без округления!
	fmt.Println(i2)
}
