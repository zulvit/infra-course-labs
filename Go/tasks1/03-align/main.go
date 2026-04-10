package main

import (
	"fmt"
	"unsafe"
)

type S1 struct {
	a bool
	b int64
	c bool
}

type S2 struct {
	a bool
	c bool
	b int64
}

func main() {
	fmt.Println("Size of S1:", unsafe.Sizeof(S1{}))
	fmt.Println("Size of S2:", unsafe.Sizeof(S2{}))
}
