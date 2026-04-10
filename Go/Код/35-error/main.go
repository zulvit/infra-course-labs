package main

import (
	"errors"
	"fmt"
)

// Кастомный тип ошибки
// Любой тип, реализующий метод Error() string, является ошибкой.
// Интерфейс error определён в стандартной библиотеке:
//
//   type error interface {
//       Error() string
//   }

// ValidationError — пример кастомной ошибки с дополнительным полем.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field %q — %s", e.Field, e.Message)
}

// Sentinel errors — заранее известные ошибки-значения.
// Объявляются на уровне пакета. Проверяются через errors.Is.
var (
	ErrNotFound   = errors.New("not found")
	ErrPermission = errors.New("permission denied")
)

func main() {
	// errors.New — создание простой ошибки
	err := errors.New("something went wrong")
	fmt.Println("err:", err) // something went wrong

	// Идиома if err != nil
	// Это повторяющийся паттерн в Go.
	// Каждый вызов, который может вернуть ошибку, проверяется сразу.
	// Нет исключений, нет try/catch — только явная проверка.
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err) // cannot divide by zero
	} else {
		fmt.Println("Результат:", result)
	}

	// fmt.Errorf с %w — оборачивание (wrapping) ошибки
	// %w сохраняет оригинальную ошибку в цепочке.
	// Это позволяет добавить контекст, не теряя информацию.
	_, err = findUser(999)
	if err != nil {
		fmt.Println("findUser:", err) // findUser: user 999: not found
	}

	// errors.Is — проверка ошибки по значению
	// Проходит по всей цепочке обёрнутых ошибок.
	if errors.Is(err, ErrNotFound) {
		fmt.Println("errors.Is: это ErrNotFound") // ✓
	}

	// errors.As — проверка ошибки по типу
	// Позволяет извлечь конкретный тип ошибки из цепочки.
	_, err = validateInput("")
	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("errors.As: field=%q, message=%q\n",
				valErr.Field, valErr.Message)
		}
	}

	// Wrapping цепочки
	// Ошибки можно оборачивать многократно, добавляя контекст на каждом уровне.
	err = processRequest("", 999)
	if err != nil {
		fmt.Println("processRequest:", err)
		// processRequest: processing request: validateInput: validation error: field "name" — must not be empty

		// Можно проверить любую ошибку в цепочке:
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Println("Глубоко в цепочке нашли ValidationError:", valErr.Field)
		}
	}

	// Правило Bill Kennedy: не логируй И не передавай
	// Ошибку нужно ЛИБО залогировать и обработать,
	// ЛИБО обернуть и вернуть наверх.
	// Никогда не делай оба действия — это дублирует информацию.
}

// safeDivide возвращает (результат, ошибку) — каноническая идиома Go.
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// findUser оборачивает sentinel error с контекстом через %w.
func findUser(id int) (string, error) {
	// Имитация поиска пользователя.
	return "", fmt.Errorf("user %d: %w", id, ErrNotFound)
}

// validateInput возвращает кастомную ошибку.
func validateInput(name string) (string, error) {
	if name == "" {
		return "", &ValidationError{
			Field:   "name",
			Message: "must not be empty",
		}
	}
	return name, nil
}

// processRequest демонстрирует многоуровневое оборачивание.
func processRequest(name string, userID int) error {
	_, err := validateInput(name)
	if err != nil {
		return fmt.Errorf("processing request: validateInput: %w", err)
	}

	_, err = findUser(userID)
	if err != nil {
		return fmt.Errorf("processing request: findUser: %w", err)
	}

	return nil
}
