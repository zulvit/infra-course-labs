package main

import "fmt"

func main() {
	// Форма 1: классическая (аналог for в C)
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0, 1, 2, 3, 4
	}

	// Форма 2: только условие (аналог while)
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n) // 128

	// // Форма 3: бесконечный цикл
	// for {
	// 	input := readInput()
	// 	if input == "quit" {
	// 		break
	// 	}
	// 	process(input)
	// }

	// i объявленная в for - видна только в теле цикла:
	for i := 0; i < 3; i++ {
	}
	// fmt.Println(i)  // ОШИБКА: i undefined

	// Несколько переменных в for:
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	// range по строке - итерация по рунам:
	for i, r := range "Hello, мир" {
		fmt.Printf("%d: %c\n", i, r)
	}

	// range по целому числу - Go 1.22+:
	for i := range 5 { // i = 0, 1, 2, 3, 4
		fmt.Println(i)
	}

	// Тип i == тип аргумента:
	n = 10
	for i := range n { // i имеет тип int64
		_ = i
	}

	// n <= 0 - цикл не выполняется, без паники:
	for range -5 {
	} // 0 итераций

	// break - выход из ближайшего цикла/switch
	// continue - следующая итерация ближайшего цикла

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} // пропустить 3
		if i == 7 {
			break
		} // выйти при 7
		fmt.Println(i)
	}
	// Вывод: 0 1 2 4 5 6

	// Метки для вложенных циклов:
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // выходим из ВНЕШНЕГО цикла
			}
			fmt.Println(i, j)
		}
	}

	// continue с меткой:
loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue loop // следующая итерация ВНЕШНЕГО цикла
			}
			fmt.Println(i, j)
		}
	}
}

// Метки использовать нежелательно, это грязный трюк
