package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "  Hello, World!  "

	// Поиск
	strings.Contains(s, "World")    // true
	strings.HasPrefix(s, "  Hello") // true
	strings.HasSuffix(s, "!  ")     // true
	strings.Index(s, "World")       // 9
	strings.Count(s, "l")           // 3
	strings.ContainsAny(s, "aeiou") // true

	// Модификация (возвращают НОВУЮ строку — строки неизменяемы!)
	strings.TrimSpace(s)                 // "Hello, World!"
	strings.Trim(s, " ")                 // "Hello, World!"
	strings.TrimLeft(s, " ")             // "Hello, World!  "
	strings.ToUpper(s)                   // "  HELLO, WORLD!  "
	strings.ToLower(s)                   // "  hello, world!  "
	strings.Replace(s, "World", "Go", 1) // "  Hello, Go!  "
	strings.ReplaceAll(s, "l", "L")      // "  HeLLo, WorLd!  "

	// Разбиение и сборка
	parts := strings.Split("a,b,c", ",") // ["a", "b", "c"]
	strings.Join(parts, " - ")           // "a - b - c"
	strings.Fields("  foo bar  baz  ")   // ["foo", "bar", "baz"]
	strings.SplitN("a,b,c,d", ",", 3)    // ["a", "b", "c,d"] — максимум N частей

	// int → string (НЕ через string(42) — это руна!)
	s1 := strconv.Itoa(42) // "42"
	// это руна с кодом 42, не "42"!
	fmt.Println(s1)

	// string → int
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(n) // 42

	// Для большего контроля:
	n64, err := strconv.ParseInt("FF", 16, 64)  // hex → int64: 255
	u64, err := strconv.ParseUint("42", 10, 64) // decimal → uint64
	fmt.Println(n64, u64)
	// float64 и string
	s2 := strconv.FormatFloat(3.14159, 'f', 2, 64) // "3.14"
	f, err := strconv.ParseFloat("3.14", 64)
	fmt.Println(s2, f)

	// bool и string
	s3 := strconv.FormatBool(true)   // "true"
	b, err := strconv.ParseBool("1") // true (принимает: "1","t","T","true","TRUE")
	fmt.Println(s3, b)
}
