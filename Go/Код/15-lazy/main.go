package main

import "fmt"

func main() {
	// && — если левая часть false, правая НЕ ВЫЧИСЛЯЕТСЯ
	// || — если левая часть true, правая НЕ ВЫЧИСЛЯЕТСЯ

	var s *string // nil

	// s != nil проверяется первым
	if s != nil && len(*s) > 0 {
		fmt.Println(*s)
	}

	// если s == nil, len(*s) вызовет панику, так как разыменование nil
	// if len(*s) > 0 && s != nil { // panic!
	// 	fmt.Println(*s)
	// }

	// Безопасный вариант с тем же смыслом проверки.
	if s == nil || len(*s) == 0 {
		fmt.Println("s is nil or empty")
	}
}
