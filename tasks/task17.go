package tasks

import (
	"fmt"
	"math/rand"
)

func binarySearch(arr []int, target int) int {
	// объявление границ поиска
	low, high := 0, len(arr)-1

	// цикл поиска пока нижняя граница поиска не больше верхней
	for low <= high {

		mid := low + (high-low)/2 // вычисление среднего индекса

		// проверка на нахождение элемента
		if arr[mid] == target {
			return mid // возвращение индекса найденного элемента в случае успеха
		}

		// сравнение элемента в середине слайса с искомым
		if arr[mid] < target {
			// если искомый элемент больше, переход к правой части слайса
			low = mid + 1
		} else {
			// если искомый элемент меньше, переход к левой части слайса
			high = mid - 1
		}
	}

	// если элемент не был найден, возвращение -1
	return -1
}

func GetTask17() {
	input := []int{-615345, -512421, -21, -13, -5, -3, -1, 0, 2, 4, 5, 6, 7, 9, 11, 412412, 516155} // объявление сортированного слайса

	// вывод изначального слайса
	fmt.Println("Initial slice:")
	fmt.Println(input)

	// выбор случайного элемента из слайса для дальнейшего поиска
	searchElement := input[rand.Intn(len(input)-1)]
	fmt.Println("Searching for the element with value:", searchElement)

	result := binarySearch(input, searchElement) // бинарный поиск элемента

	// вывод результата поиска
	if result != -1 {
		fmt.Println("Found element at index:", result)
	} else {
		fmt.Println("No element was found")
	}

	// выбор элемента для поиска, не существующего в слайсе
	searchElement = 3
	fmt.Println("Searching for the element with value:", searchElement)

	result = binarySearch(input, searchElement) // бинарный поиск элемента

	// вывод результата поиска
	if result != -1 {
		fmt.Println("Found element at index:", result)
	} else {
		fmt.Println("No element was found")
	}
}
