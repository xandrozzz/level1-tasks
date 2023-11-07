package tasks

import (
	"fmt"
)

func GetTask3() {

	arr := [5]int{2, 4, 6, 8, 10} // объявление массива чисел

	channels := make([]chan *int, 0) // объявление слайса каналов для получения данных от горутин

	for _, v := range arr {
		channel := make(chan *int) // объявление канала для получения данных от горутины
		// запуск горутины
		go func(v int) {
			result := v * v
			channel <- &result // запись результата в канал
		}(v)
		channels = append(channels, channel) // добавление канала в слайс каналов
	}

	sum := 0 // объявление переменной для суммы

	// запуск цикла для сбора данных от всех горутин по порядку
	for i := 0; i < len(channels); i++ {
		response := *<-channels[i] // получение данных из канала
		sum += response            // прибавление результата к сумме
	}

	fmt.Println(sum)
}
