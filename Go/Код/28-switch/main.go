package main

import "fmt"

func main() {
	x := 10

	switch x {
	case 1:
		fmt.Println("one")
	case 2, 3, 4: // несколько значений в одном case
		fmt.Println("two, three or four")
	case 10:
		fmt.Println("ten")
	default: // если ни один case не совпал
		fmt.Println("other")
	}

	// Выполняется ТОЛЬКО совпавший case. Нет fall-through!
	// В C нужен break; в Go break не нужен.

	// Явный fallthrough (редко нужен):
	switch x {
	case 10:
		fmt.Println("десять")
		fallthrough // явно «провалиться» в следующий case
	case 20:
		fmt.Println("двадцать или провалились")
	}

	// switch с инициализатором (как if):
	switch n := computeValue(); {
	case n < 0:
		fmt.Println("negative")
	case n == 0:
		fmt.Println("zero")
	default:
		fmt.Println("positive:", n)
	}

	score := 85

	// switch без выражения = switch true - замена цепочки if-else:
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	// default выполняется последним НЕЗАВИСИМО от расположения:
	switch x {
	default: // физически первый, но логически последний!
		fmt.Println("default")
	case 1:
		fmt.Println("one")
	}
	// Если x == 1: выведет "one"
	// Если x != 1: выведет "default"
	// Порядок case проверяется сверху вниз,
	// default - только если ничего не совпало.
}

func computeValue() int {
	return -5
}
