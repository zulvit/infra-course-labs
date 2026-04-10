package main

import (
	"fmt"
)

func main() {
	// всегда предсказуемо
	var i int            // 0
	var f float64        // 0.0
	var b bool           // false
	var s string         // "" (пустая строка)
	var p *int           // nil
	var sl []int         // nil
	var m map[string]int // nil

	fmt.Printf("i: %d, f: %f, b: %t, s: '%s', p: %v, sl: %v, m: %v\n", i, f, b, s, p, sl, m)
}
