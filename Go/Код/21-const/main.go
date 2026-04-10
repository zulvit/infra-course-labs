package main

import "fmt"

func main() {
	// Типизированная константа - жёсткий тип
	const MaxSize int = 100
	const Pi float64 = 3.14159265

	// Нетипизированная - "идеальный тип", адаптируется по контексту
	const MaxItems = 100  // "untyped int"
	const Ratio = 0.5     // "untyped float"
	const Greeting = "hi" // "untyped string"

	var x float64 = MaxItems // OK: 100 адаптируется к float64
	var y int = MaxItems     // OK: 100 адаптируется к int

	var z float64 = Pi // OK: типизированная float64
	// var w float32 = Pi // ОШИБКА: Pi типизирована как float64

	fmt.Println(x, y, z)
	fmt.Println(MaxSize, Ratio, Greeting)

	// Нетипизированные константы вычисляются с произвольной точностью:
	const BigNum = 1 << 62 // OK: вычисляется во время компиляции

	// Константы вычисляются на этапе компиляции:
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)
	fmt.Println(GB) // 1073741824
}
