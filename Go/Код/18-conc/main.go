package main

import (
	"fmt"
	"strings"
)

func main() {
	// Оператор + — удобно, но дорого в цикле
	s := "Hello" + ", " + "World" // нормально для пары строк
	fmt.Println(s)

	// В цикле — O(n²), каждый раз новая аллокация:
	var result string
	words := []string{"Go", "is", "awesome", "!"}
	for _, w := range words {
		result += w + " " // медленно!
	}

	// strings.Builder — правильный способ:

	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(w)
		sb.WriteByte(' ')
	}
	result = sb.String()

	// strings.Join — для слайса строк:
	joined := strings.Join(words, " ") // "Go is awesome !"

	fmt.Println(result)
	fmt.Println(joined)

	// Ориентировочная разница (10 000 итераций):
	// + в цикле:          ~5 ms,  50 MB аллокаций
	// strings.Builder:  ~0.1 ms, 0.1 MB аллокаций

}
