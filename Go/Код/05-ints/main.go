package main

import (
	"unsafe"
)

func main() {
	// int - ваш выбор по умолчанию для целых чисел
	var count int = 0
	for i := 0; i < 10; i++ { // i - тип int
		count++
	}

	// Почему int, а не int64 - потому что int - это "родной" тип платформы,
	// оптимизированный для производительности.
	// На 64-битных системах int обычно 64 бита, на 32-битных - 32 бита.

	// uint - для чисел без знака
	var fileSize uint = 1024 * 1024 // всегда >= 0

	// uintptr - для хранения адресов памяти (только с unsafe)
	var addr uintptr = uintptr(unsafe.Pointer(&count))

	// В учебном примере явно отмечаем использование переменных,
	// чтобы код компилировался без ошибок "declared and not used".
	_ = fileSize
	_ = addr

}
