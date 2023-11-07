package tasks

import (
	"fmt"
	"slices"
)

type sampleStruct struct {
	someData int
}

func GetTask23() {
	// объявление слайса
	sampleSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial slice:", sampleSlice)

	// получение индекса для удаления из слайса
	var index int
	fmt.Print("Enter the index of an element you want to delete: ")
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	fmt.Println("Maintaining order:")
	fmt.Println("Using slices.Delete():")

	sampleSlice = slices.Delete(sampleSlice, index, index+1) // использование функции slices.Delete() для удаления элемента из слайса

	fmt.Println("Modified slice:", sampleSlice) // вывод получившегося слайса

	fmt.Println("Using builtin methods:")

	sampleSlice = []int{1, 2, 3, 4, 5} // переопределение изначального слайса

	// удаление элемента через реслайсинг
	sampleSlice = append(sampleSlice[:index], sampleSlice[index+1:]...) // переопределение слайса на составленный из двух частей старого - до индекса удаления и после него
	fmt.Println("Modified slice:", sampleSlice)                         // вывод получившегося слайса

	fmt.Println("Not maintaining order:")

	sampleSlice = []int{1, 2, 3, 4, 5} // переопределение изначального слайса

	sampleSlice[index] = sampleSlice[len(sampleSlice)-1] // приравнивание элемента с индексом удаления к последнему
	sampleSlice = sampleSlice[:len(sampleSlice)-1]       // переопределение слайса на старый без последнего элемента

	fmt.Println("Modified slice:", sampleSlice) // вывод получившегося слайса

	fmt.Println("When deleting from slice of pointers:")

	// объявление слайса указателей на структуры
	structSlice := []*sampleStruct{{1}, {2}, {3}, {4}, {5}}
	fmt.Println("Initial slice:", structSlice)

	// получение индекса для удаления из слайса
	fmt.Print("Enter the index of an element you want to delete: ")
	_, err = fmt.Scan(&index)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	fmt.Println("Maintaining order:")
	fmt.Println("Using slices.Delete():")

	structSlice[index] = nil                                 // обнуление указателя, чтобы объект мог быть удален из памяти через GC
	structSlice = slices.Delete(structSlice, index, index+1) // использование функции slices.Delete() для удаления элемента из слайса

	fmt.Println("Modified slice:", structSlice) // вывод получившегося слайса

	fmt.Println("Using builtin methods:")

	// обнуление указателей на все элементы слайса для их удаления через GC
	for i := range structSlice {
		structSlice[i] = nil
	}
	// переопределение слайса указателей
	structSlice = []*sampleStruct{{1}, {2}, {3}, {4}, {5}}
	fmt.Println("Redefined initial slice:", structSlice)

	// удаление элемента через реслайсинг
	copy(structSlice[index:], structSlice[index+1:]) // сдвиг элементов на 1 индекс влево, так чтобы удаляемый элемент оказался в конце
	structSlice[len(structSlice)-1] = nil            // обнуление указателя, чтобы объект мог быть удален из памяти через GC
	structSlice = structSlice[:len(structSlice)-1]   // переопределение слайса на старый без последнего элемента

	fmt.Println("Modified slice:", structSlice) // вывод получившегося слайса

	fmt.Println("Not maintaining order:")

	// обнуление указателей на все элементы слайса для их удаления через GC
	for i := range structSlice {
		structSlice[i] = nil
	}
	// переопределение слайса указателей
	structSlice = []*sampleStruct{{1}, {2}, {3}, {4}, {5}}
	fmt.Println("Redefined initial slice:", structSlice)

	structSlice[index] = structSlice[len(structSlice)-1] // приравнивание элемента с индексом удаления к последнему
	structSlice[len(structSlice)-1] = nil                // обнуление указателя, чтобы объект мог быть удален из памяти через GC
	structSlice = structSlice[:len(structSlice)-1]       // переопределение слайса на старый без последнего элемента

	fmt.Println("Modified slice:", structSlice) // вывод получившегося слайса
}
