package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := make([]struct{}, 1_000_000)
	fmt.Println("Size of slice 'a':", unsafe.Sizeof(a))

	x, y := struct{}{}, struct{}{}
	fmt.Printf("%p == %p ?\n", &x, &y)
}
