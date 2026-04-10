package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int

	// Scan - читает пробелами/переносами как разделителями
	fmt.Scan(&a, &b) // ввод: "1 2\n" → a=1, b=2
	// ввод: "1\n2\n" → тоже a=1, b=2

	// Scanln - останавливается на \n
	fmt.Scanln(&a, &b) // ввод: "1 2\n" → a=1, b=2
	// ввод: "1\n2\n" → a=1, b=0 (!)

	// Scanf - строгий формат
	fmt.Scanf("%d %d", &a, &b)

	fmt.Scanf("%d", &a) // читает число, НО не поглощает \n!
	fmt.Scanln(&b)      // читает оставшийся \n как пустую строку!

	// Правильно использовать только один тип чтения,
	// или явно «съесть» \n:
	fmt.Scanf("%d\n", &a) // поглотить \n
	fmt.Scanln(&b)        // теперь работает

	// Для чтения строки с пробелами:
	var line string
	fmt.Scanln(&line) // читает до первого пробела, не всю строку
	// Для всей строки - bufio.Scanner:
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
}
