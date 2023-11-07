package tasks

import "fmt"

func GetTask12() {

	inputStrings := []string{"cat", "cat", "dog", "cat", "tree"} // входные данные

	set := make(map[string]bool) // объявление map с ключами string и значениями bool в качестве множества

	// цикл для добавления в множество элементов исходного слайса
	for _, s := range inputStrings {
		set[s] = true // добавление элемента слайса в множество
	}

	fmt.Println(set) // вывод множества

}
