package tasks

import (
	"fmt"
	"sync"
)

func GetTask2() {
	arr := [5]int{2, 4, 6, 8, 10} // объявление массива чисел

	fmt.Println("Unordered")

	var wg sync.WaitGroup // объявление WaitGroup для обеспечение выполнения всех горутин

	for _, v := range arr {
		wg.Add(1) // увеличение счетчика WaitGroup
		// запуск горутины
		go func(v int, wg *sync.WaitGroup) {
			fmt.Println(v * v)
			wg.Done() // уменьшение счетчика WaitGroup
		}(v, &wg)
	}

	wg.Wait() // ожидание выполнения всех горутин

	fmt.Println("Ordered")

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

	// запуск цикла для сбора данных от всех горутин по порядку
	for i := 0; i < len(channels); i++ {
		response := *<-channels[i] // получение данных из канала
		fmt.Println(response)      // вывод данных
	}

}
