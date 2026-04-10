package main

import "demos/calc"

func main() {
	// Вызов функции Calculate из пакета calc
	result := calc.Calculate(3, 5)
	println("Результат:", result)
}
