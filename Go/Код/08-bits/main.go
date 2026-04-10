package main

import (
	"fmt"
)

func main() {
	// Операторы:
	// &   — AND (И)
	// |   — OR (ИЛИ)
	// ^   — XOR (бинарный) / NOT (унарный — в Go нет ~, используют ^)
	// &^  — AND NOT (bit clear)
	// <<  — сдвиг влево
	// >>  — сдвиг вправо

	a := 0b1010 // 10
	b := 0b1100 // 12

	fmt.Printf("%04b\n", a&b)  // 1000 = 8  (AND)
	fmt.Printf("%04b\n", a|b)  // 1110 = 14 (OR)
	fmt.Printf("%04b\n", a^b)  // 0110 = 6  (XOR)
	fmt.Printf("%04b\n", a&^b) // 0010 = 2  (AND NOT)
	fmt.Printf("%04b\n", a<<1) // 10100 = 20 (сдвиг = *2)
	fmt.Printf("%04b\n", a>>1) // 0101 = 5  (сдвиг = /2)

	// Унарный ^ = побитовое НЕ (NOT):
	x := 5
	fmt.Printf("%b\n", ^x) // -6 - инвертирует все биты
	// В Go нет оператора ~, вместо него используется унарный ^ для побитового NOT.

	// Битовые флаги:
	const (
		Read    = 1 << iota // 001
		Write               // 010
		Execute             // 100
	)
	perm := Read | Write
	fmt.Println(perm&Execute != 0) // false
}
