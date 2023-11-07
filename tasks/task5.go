package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GetTask5() {
	channel := make(chan interface{}) // создание канала для любых данных

	var n int
	// получение длины таймаута
	fmt.Print("Enter the timeout in seconds: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second) // создание контекста с таймаутом и объявление функции завершения
	defer cancel()                                                                         // вызов функции завершения в defer, чтобы избежать утечки контекста

	var wg sync.WaitGroup // объявление WaitGroup для обеспечение выполнения всех горутин

	wg.Add(1) // увеличение счетчика WaitGroup
	go func(ctx context.Context, wg *sync.WaitGroup) {
		for {
			// неблокирующее получение данных из канала
			select {
			case data := <-channel:
				fmt.Println(data)
			default:
			}

			// неблокирующая проверка контекста на завершение
			select {
			case <-ctx.Done():
				fmt.Println("Receiver stopped")
				wg.Done() // уменьшение счетчика WaitGroup
				return
			default:
			}
		}

	}(ctx, &wg)

	wg.Add(1) // увеличение счетчика WaitGroup
	// запуск горутины для записи данных в канал
	go func(ctx context.Context, wg *sync.WaitGroup) {
		for {
			// запись случайных данных в канал
			x := rand.Intn(5)
			switch x {
			case 0:
				channel <- "aboba"
			case 1:
				channel <- 9412412
			case 2:
				channel <- []int{1, 2, 5, 5, 94129, 0, -2}
			case 3:
				channel <- true
			case 4:
				channel <- 21.3
			}
			time.Sleep(500 * time.Millisecond)

			// неблокирующая проверка контекста на завершение
			select {
			case <-ctx.Done():
				fmt.Println("Sender stopped")
				wg.Done() // уменьшение счетчика WaitGroup
				return
			default:
			}
		}

	}(ctx, &wg)

	wg.Wait() // ожидание завершения всех горутин
	fmt.Println("All goroutines stopped, exiting")
}
