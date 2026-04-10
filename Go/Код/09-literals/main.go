package main

import (
	"fmt"
)

func main() {
	// Десятичные (обычные)
	x := 1000000 // трудно читать
	fmt.Printf("%d\n", x)

	x = 1_000_000 // разделитель _ для читаемости

	// Двоичные (binary) - префикс 0b
	bin := 0b1010_1010 // = 170
	fmt.Printf("%d\n", bin)

	// Восьмеричные (octal) - префикс 0o
	oct := 0o755            // права доступа Unix: rwxr-xr-x
	old := 0755             // старый синтаксис (тоже восьмеричный, но менее явный)
	fmt.Printf("%d\n", oct) // 493
	fmt.Printf("%d\n", old) // 493

	// Шестнадцатеричные (hex) - префикс 0x
	hex := 0xFF_FF_FF // цвет в RGB
	ip := 0xC0A80001  // 192.168.0.1
	fmt.Printf("%d\n", hex)
	fmt.Printf("%x\n", ip) // c0a80001

	// Числа с плавающей точкой
	pi := 1_000.555_555
	sci := 6.022e23 // научная нотация: 6.022 × 10²³

	// Комплексные литералы:
	c := 3 + 4i // complex128 литерал
	c2 := complex64(1 + 2i)

	fmt.Printf("pi: %f, sci: %e, c: %v, c2: %v\n", pi, sci, c, c2)
}
