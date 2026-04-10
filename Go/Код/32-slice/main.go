package main

import "fmt"

func main() {
	// Слайс — динамический массив в Go.
	// В отличие от массива, длина слайса НЕ является частью типа.
	// Слайс — это дескриптор из трёх полей (src/runtime/slice.go):
	//
	//   type slice struct {
	//       array unsafe.Pointer  // указатель на backing array
	//       len   int             // текущая длина
	//       cap   int             // вместимость (capacity)
	//   }
	//
	// Все операции со слайсом работают через этот дескриптор.

	// Создание через литерал
	nums := []int{1, 2, 3}
	fmt.Println("nums:", nums)                           // [1 2 3]
	fmt.Printf("len=%d, cap=%d\n", len(nums), cap(nums)) // len=3, cap=3

	// Nil slice vs empty slice
	var nilSlice []int        // nil slice: len=0, cap=0
	emptySlice := []int{}     // empty slice: len=0, cap=0, но не nil
	fmt.Println("nilSlice == nil:", nilSlice == nil)     // true
	fmt.Println("emptySlice == nil:", emptySlice == nil) // false
	// Оба имеют len=0, cap=0, но nil-слайс не аллоцирован.
	// В JSON nil → null, empty → [].

	// Чтение и запись — как в массиве
	s := []int{1, 2, 3}
	fmt.Println("s[2]:", s[2]) // 3
	s[0] = 10
	fmt.Println("s:", s) // [10 2 3]

	// Sub-slicing (срез среза)
	// Оператор : создаёт новый слайс, ссылающийся на ТОТ ЖЕ backing array.
	fmt.Println("s[1:3]:", s[1:3]) // [2 3]
	fmt.Println("s[:2]: ", s[:2])  // [10 2]
	fmt.Println("s[2:]: ", s[2:])  // [3]

	// ЛОВУШКА: shared backing array
	// Два слайса могут ссылаться на один и тот же массив.
	// Изменение через один слайс видно в другом!
	words := []string{"e", "l", "e", "m", "e", "n", "t"}
	s1 := words[:4]  // [e l e m]
	s2 := words[1:6] // [l e m e n]
	fmt.Println("s1:", s1, "s2:", s2)

	s1[1] = "L" // Меняем через s1 - видно в words и s2!
	fmt.Println("words после s1[1]=\"L\":", words) // [e L e m e n t]
	fmt.Println("s2 после изменения:   ", s2)      // [L e m e n]

	// append() - добавление элементов
	// append возвращает НОВЫЙ слайс.
	// Если cap достаточно - данные добавляются в существующий массив.
	// Если cap недостаточно - создаётся НОВЫЙ массив (с увеличенной cap).
	items := []string{"hello"}
	items = append(items, "world")
	fmt.Println("items:", items) // [hello world]

	// Добавление слайса в слайс через оператор ...:
	p1 := []int{1, 2, 3}
	p2 := []int{4, 5, 6}
	p1 = append(p1, p2...)
	fmt.Println("p1:", p1) // [1 2 3 4 5 6]

	// make() - предаллокация для эффективности
	// make([]T, len, cap) создаёт слайс с заданной длиной и вместимостью.
	m1 := make([]int, 5, 5) // len=5, заполнено zero values
	fmt.Println("make(5,5):", m1) // [0 0 0 0 0]

	m2 := make([]int, 0, 5) // len=0, но cap=5 - пустой, но преаллоцирован
	fmt.Println("make(0,5):", m2) // []
	// Использование make([]T, 0, expected) + append - идиоматичный путь,
	// когда знаешь примерный размер заранее.

	// Слайсы и цикл for
	acc := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		acc = append(acc, i)
	}
	fmt.Println("acc:", acc) // [0 1 2 3 4 5 6 7 8 9]

	// range по слайсу
	// range возвращает (индекс, значение) на каждом шаге.
	names := []string{"John", "Harold", "Vince"}
	for i, name := range names {
		fmt.Printf("  Hello %s at index %d\n", name, i)
	}

	// Передача слайса в функцию
	// Дескриптор (len, cap) копируется, но backing array - по ссылке.
	// Результат: изменение существующих элементов видно в оригинале,
	// но append внутри функции НЕ затрагивает оригинальный слайс.
	data := []int{1, 2, 3, 4, 5}
	modifySlice(data)
	fmt.Println("data после modifySlice:", data) // [1 2 10 4 5]
	// Элемент data[2] изменился на 10, но 6 не добавился.
}

// modifySlice демонстрирует поведение слайса при передаче в функцию.
// Изменение элемента - видно в оригинале (shared backing array).
// Append - НЕ видно (копия дескриптора, len оригинала не меняется).
func modifySlice(nums []int) {
	nums[2] = 10             // видно в оригинале
	nums = append(nums, 6)   // НЕ видно в оригинале
	_ = nums
}
