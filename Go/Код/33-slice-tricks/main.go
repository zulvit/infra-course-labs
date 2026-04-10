package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// copy() - безопасное копирование слайса
	// func copy(dst, src []Type) int
	// Копирует min(len(dst), len(src)) элементов.
	// dst должен иметь достаточный len (не cap, а len!).

	nums := []int{1, 2, 3, 4, 5}
	numsCp := make([]int, len(nums)) // len = 5 - правильно
	copied := copy(numsCp, nums)
	fmt.Println("copied:", copied)   // 5
	fmt.Println("numsCp:", numsCp)   // [1 2 3 4 5]

	// ЛОВУШКА: copy в слайс с len=0
	// Ничего не скопируется, потому что len(dst) == 0.
	badCopy := make([]int, 0) // len = 0!
	copy(badCopy, nums)
	fmt.Println("badCopy:", badCopy) // [] - пусто!

	// Почему нельзя просто присвоить?
	// При присваивании копируется только дескриптор,
	// но backing array остаётся общим.
	alias := nums
	alias[0] = 999
	fmt.Println("nums после alias[0]=999:", nums) // [999 2 3 4 5] - изменился!

	// Three-index slice: s[lo:hi:max]
	// Третий индекс ограничивает cap нового слайса.
	// Используется для защиты от случайного append в shared array.
	original := []int{1, 2, 3, 4, 5}
	limited := original[1:3:3] // len=2, cap=2 (а не 4)
	fmt.Printf("limited: %v, len=%d, cap=%d\n", limited, len(limited), cap(limited))
	// Теперь append на limited создаст НОВЫЙ backing array,
	// не затрагивая original.
	limited = append(limited, 99)
	fmt.Println("original после append на limited:", original) // [1 2 3 4 5] - не изменился

	// Сортировка: sort.Slice
	// sort.Slice использует quicksort (нестабильная).
	// Принимает interface{} (любой слайс) и компаратор.
	unsorted := []int{2, 1, 6, 5, 3, 4}
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i] < unsorted[j]
	})
	fmt.Println("sort.Slice:", unsorted) // [1 2 3 4 5 6]

	// sort.SliceStable - стабильная сортировка
	// Сохраняет порядок равных элементов (insertion sort).
	// Выбор между Slice и SliceStable зависит от данных:
	// SliceStable - когда порядок равных элементов важен.
	data := []int{5, 3, 1, 4, 2, 6}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	fmt.Println("sort.SliceStable:", data) // [1 2 3 4 5 6]

	// Пакет slices (Go 1.21+) - современный подход
	// Типобезопасные функции без interface{} и без компараторов для простых случаев.

	// slices.Sort - дженерик-сортировка (быстрее sort.Slice для примитивов):
	vals := []int{9, 3, 7, 1, 5}
	slices.Sort(vals)
	fmt.Println("slices.Sort:", vals) // [1 3 5 7 9]

	// slices.Contains - проверка наличия элемента:
	fmt.Println("contains 7:", slices.Contains(vals, 7)) // true
	fmt.Println("contains 2:", slices.Contains(vals, 2)) // false

	// slices.Index - индекс первого вхождения:
	fmt.Println("index of 5:", slices.Index(vals, 5)) // 2

	// slices.Reverse - разворот на месте:
	slices.Reverse(vals)
	fmt.Println("reversed:", vals) // [9 7 5 3 1]

	// Удаление элемента без утечки памяти
	// Классический приём: сдвиг через copy + обнуление последнего.
	del := []string{"a", "b", "c", "d", "e"}
	idx := 2 // удаляем "c"
	copy(del[idx:], del[idx+1:])
	del[len(del)-1] = "" // обнуляем, чтобы GC мог собрать
	del = del[:len(del)-1]
	fmt.Println("после удаления:", del) // [a b d e]

	// Или через slices.Delete (Go 1.21+):
	del2 := []string{"a", "b", "c", "d", "e"}
	del2 = slices.Delete(del2, 2, 3) // удаляем индексы [2, 3)
	fmt.Println("slices.Delete:", del2) // [a b d e]
}
