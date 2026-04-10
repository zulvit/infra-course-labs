package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 10
	if x > 0 { // () скобки вокруг условия НЕ нужны
		fmt.Println("positive")
	} else if x < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}

	// Фигурные скобки ОБЯЗАТЕЛЬНЫ даже для одной строки:
	// if x > 0   // ОШИБКА
	//     ...

	// Открывающая { ДОЛЖНА быть на той же строке:
	// if x > 0
	// {              // ОШИБКА КОМПИЛЯЦИИ
	//     ...
	// }
	// Go автоматически вставляет ; после строки, которая
	// заканчивается на идентификатор, число, строку, ), ], }

	// Синтаксис: if инициализатор; условие { }

	// Без инициализатора - переменные засоряют область видимости:
	result, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("parsed:", result)
	// result и err видны здесь, хотя они уже не нужны

	// С инициализатором - переменные ограничены блоком if:
	if result, err := strconv.Atoi("42"); err != nil {
		fmt.Println("error:", err)
		return
	} else {
		fmt.Println("parsed:", result) // result доступен только здесь
	}
	// result и err НЕ видны здесь, они были созданы внутри блока if

	s := "Hello, World!"
	if n := len(s); n == 0 {
		fmt.Println("empty")
	} else if n < 10 {
		fmt.Println("short:", n)
	} else {
		fmt.Println("long:", n)
	}
	// n недоступен снаружи блока if
}
