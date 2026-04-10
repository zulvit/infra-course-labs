package main

import "fmt"

type Engine struct{ Color string }

type Car struct {
	Engine
	Color string
}

func main() {
	c := Car{
		Color:  "Red",
		Engine: Engine{Color: "Blue"},
	}
	fmt.Println(c.Color)
}
