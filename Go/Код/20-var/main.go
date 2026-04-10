package main

import (
	"fmt"
	"strconv"
)

// := - объявляет И инициализирует. Только внутри функции!
// На уровне пакета := НЕ работает (только var).
// test := 1 //syntax error: non-declaration statement outside function body

var (
	version    = "1.0.0"
	buildTime  = "2025-01-01"
	maxRetries = 3
	myMap      = map[string]int{"key": 1}

	// Флаг для демонстрации области видимости в process.
	someCondition = true
)

func main() {
	// Полный синтаксис: var имя тип = значение
	// Скобки тут создают локальный блок для демонстрации и чтобы не было конфликта имен
	{
		var x int = 42
		var name string = "Alice"
		var active bool = true
		fmt.Println(x, name, active)
	}

	// Тип можно опустить - инферируется из значения:
	{
		var x = 42         // int
		var name = "Alice" // string
		fmt.Println(x, name)
	}

	// Значение можно опустить - нулевое значение:
	{
		var x int       // 0
		var name string // ""
		var active bool // false
		fmt.Println(x, name, active)
	}

	// Групповое объявление:
	host, port, timeout := "localhost", 8080, 30
	fmt.Println(host, port, timeout)

	// На уровне пакета (не внутри функции) - см. переменные выше.
	fmt.Println(version, buildTime, maxRetries)

	// := - объявляет И инициализирует. Только внутри функции!
	x := 42
	name := "Alice"
	x, y := 1, 2 // можно несколько
	fmt.Println(x, name, y)

	// Важный нюанс - при := хотя бы одна переменная должна быть НОВОЙ!
	// Если обе переменные уже объявлены, то используем =
	x = 10 // здесь просто присваивание уже существующей x
	// x := 20      // ОШИБКА: x уже объявлена
	x, y = 30, 40 // обе уже объявлены, поэтому используем =
	fmt.Println(x, y)

	// Частый паттерн с ошибками:
	result, err := doSomething() // OK: обе новые
	result2, err := doOther()    // OK: result2 новая, err переиспользуется
	fmt.Println(result, result2, err)

	// На уровне пакета := НЕ работает (только var). Но здесь мы ВНУТРИ функции!
	// Поэтому := создаёт НОВУЮ локальную version, затеняя (shadowing) пакетную.
	version := "1.0.0"
	fmt.Println(version)

	// Blank identifier _ - явное игнорирование значения:
	value, _ := strconv.Atoi("42") // игнорируем ошибку (осторожно!)
	_, exists := myMap["key"]      // только булевый результат
	fmt.Println(value, exists)

	x = 10
	fmt.Println(x) // 10

	if true {
		x := 20        // НОВАЯ переменная x, не присвоение!
		fmt.Println(x) // 20 - локальная x
	}

	fmt.Println(x) // 10 - внешняя x не изменилась!

}

// Частая ошибка с err:
func process() error {
	result, err := step1()
	if err != nil {
		return err
	}

	if someCondition {
		// := создаёт НОВУЮ err в блоке if - тень!
		result2, err := step2()
		if err != nil {
			return err
		} // возвращает локальную err
		_ = result2
	}
	// Внешняя err здесь - от step1, а не step2!
	_ = result
	return nil
}

func doSomething() (int, error) { return 1, nil }
func doOther() (int, error)     { return 2, nil }
func step1() (int, error)       { return 10, nil }
func step2() (int, error)       { return 20, nil }
