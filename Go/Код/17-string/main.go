package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, мир"

	// Строки — НЕИЗМЕНЯЕМЫ (immutable)
	// s[0] = 'h'   // ОШИБКА КОМПИЛЯЦИИ: cannot assign to s[0]

	// Но можно прочитать байт по индексу:
	fmt.Printf("%d %c\n", s[0], s[0]) // 72 H

	// Длина в БАЙТАХ, не в символах:
	fmt.Println(len(s)) // 13, не 10!
	// Интерпретируемые строки — двойные кавычки
	s1 := "Hello\nWorld"    // \n — символ новой строки
	s2 := "Tab:\there"      // \t — символ табуляции
	s3 := "Quote: \""       // \" — кавычка внутри
	s4 := "Unicode: \u0041" // \u0041 = 'A'
	fmt.Println(s1)
	// Hello
	// World

	fmt.Println(s2) // Tab:    here
	fmt.Println(s3)
	fmt.Println(s4) // Unicode: A

	// Сырые строки — обратные кавычки (backtick)
	// Никаких escape-последовательностей!
	s5 := `Hello\nWorld` // \n — буквально \n, не перенос!
	s6 := `Многострочная
строка
без escape`

	fmt.Println(s5) // Hello\nWorld
	fmt.Println(s6) // три строки

	// Когда использовать raw strings:
	re := `\d+\.\d+`                         // регулярные выражения
	q := `SELECT * FROM users WHERE id = $1` // SQL запросы
	fmt.Println(re)
	fmt.Println(q)
	// Go хранит строки в UTF-8.
	// ASCII-символы: 1 байт.
	// Кириллица: 2 байта. Китайский: 3 байта. Emoji: 4 байта.

	s11 := "Hello"  // 5 байт, 5 символов
	s22 := "Привет" // 12 байт, 6 символов (2 байта × 6)
	s33 := "你好"     // 6 байт, 2 символа  (3 байта × 2)

	fmt.Println(len(s11), len(s22), len(s33)) // 5 12 6

	// Правильный подсчёт символов:
	fmt.Println(utf8.RuneCountInString(s11)) // 5
	fmt.Println(utf8.RuneCountInString(s22)) // 6
	fmt.Println(utf8.RuneCountInString(s33)) // 2

	// Проверка валидности UTF-8:
	fmt.Println(utf8.ValidString(s22))        // true
	fmt.Println(utf8.ValidString("\xff\xfe")) // false
	s7 := "Привет"

	// Индексирование через [] → БАЙТ (uint8)
	fmt.Println(s7[0])        // 208 — первый байт буквы П
	fmt.Printf("%c\n", s7[0]) // Ð — неправильный символ!

	// Для получения руны — конвертировать в []rune:
	runes := []rune(s7)          // []rune{'П', 'р', 'и', 'в', 'е', 'т'}
	fmt.Printf("%c\n", runes[0]) // П — правильно!
	fmt.Println(len(runes))      // 6

	// utf8.DecodeRuneInString — без конвертации всей строки:
	r, size := utf8.DecodeRuneInString(s7)
	fmt.Printf("rune: %c, size: %d bytes\n", r, size)
	// rune: П, size: 2 bytes

	// Срез строки — тоже по байтам!
	s8 := s7[0:4] // первые 4 байта = 2 символа кириллицы = "Пр"

	fmt.Println(s8) // Пр

	// Безопасный срез по рунам:
	s9 := string([]rune(s7)[0:3]) // первые 3 символа
	fmt.Println(s9)               // "При"
}
