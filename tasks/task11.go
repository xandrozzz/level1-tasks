package tasks

import "fmt"

func GetTask11() {

	// объявление двух map с ключами string и значениями bool в качестве множеств
	set1 := map[string]bool{"a": true, "b": true, "c": true, "d": true}
	set2 := map[string]bool{"b": true, "d": true, "f": true, "h": true}

	intersection := make(map[string]bool) // объявление аналогичной map для пересечения

	// цикл для нахождения элементов, находящихся в обоих множествах
	for k, _ := range set1 {
		// проверка нахождения ключа в другом множестве
		if set2[k] {
			intersection[k] = true // добавление ключа в пересечение
		}
	}

	fmt.Println(intersection) // вывод множества

}
