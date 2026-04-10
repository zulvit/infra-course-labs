package main

import "fmt"

type A struct {
	ID    int
	Flags [2]bool
}

type B struct {
	ID    int
	Flags []bool
}

func main() {
	a1, a2 := A{1, [2]bool{true}}, A{1, [2]bool{true}}
	b1, b2 := B{1, []bool{true}}, B{1, []bool{true}}

	fmt.Println(a1 == a2)
	fmt.Println(b1 == b2)
}
