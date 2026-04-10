package main

import "fmt"

func main() {
	// byte = uint8 - одиночный байт
	var b byte = 'A'       // ASCII код: 65
	fmt.Println(b)         // 65
	fmt.Println(string(b)) // "A"

	// rune = int32 - Unicode code point
	var r rune = 'А'       // Кириллическая А, U+0410: 1040
	fmt.Println(r)         // 1040
	fmt.Println(string(r)) // "А"

	// Разница byte vs rune на строках:
	s := "Hello, мир"
	fmt.Println(len(s))         // 13 байт а не 13 символов
	fmt.Println([]byte(s))      // [72 101 108 108 111 44 32 208 188...]
	fmt.Println([]rune(s))      // [72 101 108 108 111 44 32 1084 1080 1088]
	fmt.Println(len([]rune(s))) // 10 символов

	// byte - для работы с сырыми данными, сетью, файлами
	// rune - для работы с текстом, Unicode
}
