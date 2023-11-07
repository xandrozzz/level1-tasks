package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func GetTask4() {

	channel := make(chan interface{}) // создание канала для любых данных

	var n int
	// получение количества воркеров
	fmt.Print("Enter the number of workers: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	var wg sync.WaitGroup // объявление WaitGroup для обеспечение выполнения всех горутин

	ctx, cancel := context.WithCancel(context.Background()) // создание контекста и объявление функции завершения

	// создание воркеров
	for i := 0; i < n; i++ {

		wg.Add(1) // увеличение счетчика WaitGroup
		// запуск горутины воркера
		go func(wn int, ctx context.Context, wg *sync.WaitGroup) {
			for {
				// неблокирующее получение данных из канала
				select {
				case data := <-channel:
					fmt.Println("From worker", wn, ":", data)
				default:
				}

				// неблокирующая проверка контекста на завершение
				select {
				case <-ctx.Done():
					fmt.Println("Worker", wn, "stopped")
					wg.Done() // уменьшение счетчика WaitGroup
					return
				default:
				}
			}

		}(i, ctx, &wg)
	}

	// запуск горутины для записи данных в канал
	wg.Add(1) // увеличение счетчика WaitGroup
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

	signalCh := make(chan os.Signal, 1)   // создание канала для отслеживание сигнала о прерывании
	signal.Notify(signalCh, os.Interrupt) // направление сигнала о прерывании в канал

	// запуск горутины для завершения контекста при получении сигнала о прерывании
	go func() {
		select {
		case <-signalCh:
			cancel() // завершение контекста
			return
		}
	}()

	wg.Wait() // ожидание завершения всех горутин
	fmt.Println("All goroutines stopped, exiting")

}
