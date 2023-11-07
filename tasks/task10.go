package tasks

import "fmt"

func GetTask10() {
	groups := make(map[int][]float32) // создание map с ключами int и значениями []float32

	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // входные данные
	fmt.Println("Input data:", temperature)

	for _, temp := range temperature {
		index := (int(temp) / 10) * 10              // определение ключа map исходя из температуры
		groups[index] = append(groups[index], temp) // добавление температуры в соответствующий слайс внутри map
	}

	// вывод групп и их ключей
	fmt.Println("Result groups:")
	for key, group := range groups {
		fmt.Println(key, group)
	}

}
