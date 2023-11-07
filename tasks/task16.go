package tasks

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	newArr := make([]int, len(arr)) // создаем слайс, в который будет записан результат

	copy(newArr, arr) // копируем слайс из параметра функции в результирующий, чтобы не изменять изначальный слайс

	sort(newArr, 0, len(arr)-1) // вызывает функцию sort на целый слайс

	return newArr
}

func sort(arr []int, start, end int) {

	// проверка на размер слайса, выход, если он слишком маленький
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]   // выбираем поворотный элемент (можно взять случайное значение из слайса)
	splitIndex := start // объявление переменной для хранения индекса, по которому произойдет разделение слайса, изначально индекс равен стартовому для рассматриваемой части слайса

	// цикл для частичной сортировки части слайса
	for i := start; i < end; i++ {
		// проверка, меньше ли элемент чем поворотный элемент
		if arr[i] < pivot {

			arr[splitIndex], arr[i] = arr[i], arr[splitIndex] // если элемент меньше поворотного, он перемещается в начало части слайса

			splitIndex++ // начало части слайса сдвигается, так как перед ним находятся уже сортированые элементы
		}
	}

	// перемещаем поворотный элемент на позицию перед последним элементом, меньшим чем тот, который находится между получившимися двумя частями слайса
	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	// рекурсивный вызов функции sort для двух получившихся частей слайса
	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}

func generateSlice(size int) []int {
	// объявление слайса
	slice := make([]int, size)
	// заполнение слайса случайными числами
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func GetTask16() {
	input := generateSlice(50) // генерация слайса из случайных чисел указанной длины

	// вывод изначального слайса
	fmt.Println("Initial slice:")
	fmt.Println(input)

	input = quickSort(input) // сортировка слайса алгоритом quicksort

	// вывод отсортированного слайса
	fmt.Println("Sorted slice:")
	fmt.Println(input)
}
