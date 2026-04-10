package main

import "fmt"

// iota сбрасывается в 0 для каждой группы const
type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

// Пропуск значения через _:
type Weekday int

const (
	_         = iota // 0 — пропускаем
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
	Sunday           // 7
)

// Битовые маски:
type Permission uint

const (
	Read    Permission = 1 << iota // 1  = 001
	Write                          // 2  = 010
	Execute                        // 4  = 100
)

func main() {
	p := Read | Write // побитовое ИЛИ - 001 | 010 = 011 (= 3)
	fmt.Println(p&Execute != 0) // побитовое И - 011 & 100 = 000 (= 0) => false
}
