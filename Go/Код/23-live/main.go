package main

import "fmt"

var packageLevel = "виден во всём пакете" // уровень пакета

func myFunc() {
	funcLevel := "виден в функции"
	condition := true

	if condition {
		blockLevel := "виден в блоке if"
		fmt.Println(packageLevel) // OK: виден во всём пакете
		fmt.Println(funcLevel)
		fmt.Println(blockLevel)
		_ = funcLevel // OK: внешняя область доступна
	}
	// _ = blockLevel  // ОШИБКА: blockLevel недоступен	
}

func main() {
	myFunc()
	fmt.Println(MaxRetries)
	GetUser()
}

// Правила именования:
// camelCase - для всего внутри пакета
// userAge := 25
// maxRetryCount := 3

// PascalCase — для экспортируемого
func GetUser() {}

const MaxRetries = 3

// Аббревиатуры — все буквы в одном регистре:
// userID  := 1     // не userId
// parseURL := ""   // не parseUrl
// var httpClient   // не HttpClient (внутри пакета)

// Однобуквенные — только для коротких областей:
// for i := range items { }
// for k, v := range m { }
