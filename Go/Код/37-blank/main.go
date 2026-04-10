package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Оператор _ (blank identifier)
	// В Go каждая объявленная переменная ДОЛЖНА быть использована.
	// Компилятор выдаст ошибку, если переменная не используется.
	// Blank identifier _ позволяет явно игнорировать значения.

	// 1. Игнорирование возвратов функции
	// Когда функция возвращает несколько значений,
	// но нужны не все из них.

	// Игнорируем индекс, берём только значение:
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		fmt.Println(name)
	}

	// Игнорируем значение, берём только индекс:
	for i := range names {
		fmt.Println("index:", i)
	}

	// Игнорируем оба — range только для подсчёта итераций:
	count := 0
	for range names {
		count++
	}
	fmt.Println("count:", count) // 3

	// 2. Игнорирование ошибки
	// Иногда ошибку игнорируют осознанно.
	// Это допустимо ТОЛЬКО когда ошибка действительно неважна.
	fmt.Fprintf(os.Stdout, "hello\n") // игнорируем int (количество байт)

	// 3. Import for side-effects
	// Иногда пакет нужен не для его функций,
	// а для побочных эффектов его init().
	// Пример: import _ "net/http/pprof" — регистрирует обработчики профилирования.
	// Пример: import _ "image/png" — регистрирует PNG-декодер.
	//
	// import _ "net/http/pprof"
	//
	// Без _ компилятор выдаст ошибку "imported and not used".

	// 4. Compile-time interface check
	// Проверяем, что тип реализует интерфейс,
	// без создания реальной переменной.
	// Если myWriter НЕ реализует io.Writer — ошибка компиляции.
	var _ io.Writer = (*myWriter)(nil)

	// Это идиоматичный способ гарантировать совместимость типа
	// с интерфейсом на этапе компиляции.

	// 5. Множественные возвраты с _
	// Когда функция возвращает (результат, ошибка),
	// а нам нужна только ошибка для проверки:
	_, err := fmt.Println("test")
	if err != nil {
		fmt.Println("ошибка записи:", err)
	}

	// Использование myWriter
	w := myWriter{}
	_, _ = w.Write([]byte("Go is great!")) // Go is great!
}

// myWriter — пример типа, реализующего io.Writer.
type myWriter struct{}

func (w *myWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	fmt.Println()
	return len(p), nil
}
