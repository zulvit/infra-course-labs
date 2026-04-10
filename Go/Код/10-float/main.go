package main

import (
	"fmt"
)

func main() {
	var f32 float32 = 1.23456789
	var f64 float64 = 1.23456789

	fmt.Printf("float32: %.10f\n", f32) // 1.2345678806 - потеря!
	fmt.Printf("float64: %.10f\n", f64) // 1.2345678900 - точно

	// используйте float64 по умолчанию
	// float32 - только при явной экономии памяти (большие массивы, GPU)
}
