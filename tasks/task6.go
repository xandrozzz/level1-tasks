package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sampleGoroutine1(stopChannel chan bool, wg *sync.WaitGroup) {
	for {
		// неблокирующее получение данных из канала
		select {
		case <-stopChannel:
			fmt.Println("Stopping the goroutine from example 1")
			wg.Done() // уменьшение счетчика WaitGroup
			return
		default:
			fmt.Println("Goroutine from example 1 is still running")
			time.Sleep(time.Second)
		}
	}
}

func sampleGoroutine2(stopFlag *bool, wg *sync.WaitGroup) {

	for {
		// постоянная проверка логической переменной по ссылке
		if *stopFlag {
			fmt.Println("Stopping the goroutine from example 21")
			wg.Done() // уменьшение счетчика WaitGroup
			return
		}
		fmt.Println("Goroutine from example 2 is still running")
		time.Sleep(time.Second)
	}
}

func sampleGoroutine3(stopChannel chan interface{}, wg *sync.WaitGroup) {
	for {
		// неблокирующее получение данных из канала, также сработает при закрытии канала
		select {
		case <-stopChannel:
			fmt.Println("Stopping the goroutine from example 3")
			wg.Done() // уменьшение счетчика WaitGroup
			return
		default:
			fmt.Println("Goroutine from example 3 is still running")
			time.Sleep(time.Second)
		}
	}
}

func sampleGoroutine4(ctx context.Context, wg *sync.WaitGroup) {
	for {
		// неблокирующая проверка на завершение контекста
		select {
		case <-ctx.Done():
			fmt.Println("Stopping the goroutine from example 4")
			wg.Done() // уменьшение счетчика WaitGroup
			return
		default:
			fmt.Println("Goroutine from example 4 is still running")
			time.Sleep(time.Second)
		}
	}
}

func GetTask6() {
	var wg sync.WaitGroup // объявление WaitGroup для обеспечения завершения горутин

	fmt.Println("Using channels:")

	stopChannel := make(chan bool) // объявление канала для контроля работы горутины

	wg.Add(1) // увеличение счетчика WaitGroup
	fmt.Println("Starting the goroutine")
	go sampleGoroutine1(stopChannel, &wg) // запуск горутины

	time.Sleep(3 * time.Second)

	stopChannel <- true // отправление данных в канал в качестве сигнала для остановки
	wg.Wait()           // ожидание остановки горутины

	fmt.Println("Using variables:")

	var stopFlag bool // объявление переменной

	wg.Add(1) // увеличение счетчика WaitGroup
	fmt.Println("Starting the goroutine")
	go sampleGoroutine2(&stopFlag, &wg) // запуск горутины

	time.Sleep(3 * time.Second)

	stopFlag = true // изменение значения переменной в качестве сигнала для остановки
	wg.Wait()       // ожидание остановки горутины

	fmt.Println("Using channel closure:")

	stopCloseChannel := make(chan interface{}) // объявление канала для контроля работы горутины

	wg.Add(1) // увеличение счетчика WaitGroup
	fmt.Println("Starting the goroutine")
	go sampleGoroutine3(stopCloseChannel, &wg) // запуск горутины

	time.Sleep(3 * time.Second)

	close(stopCloseChannel) // закрытие канала в качестве сигнала для остановки
	wg.Wait()               // ожидание остановки горутины

	fmt.Println("Using context:")

	ctx, cancel := context.WithCancel(context.Background()) // создание контекста для контроля работы горутины

	wg.Add(1) // увеличение счетчика WaitGroup
	fmt.Println("Starting the goroutine")
	go sampleGoroutine4(ctx, &wg) // запуск горутины

	time.Sleep(3 * time.Second)

	cancel()  // вызов функции завершения контекста в качестве сигнала для остановки
	wg.Wait() // ожидание остановки горутины
}
