package tasks

import (
	"context"
	"fmt"
	"sync"
)

func GetTask9() {
	inputChannel := make(chan int)  // создание канала для входных данных
	outputChannel := make(chan int) // создание канала для выходных данных

	inputCtx, inputCancel := context.WithCancel(context.Background())   // создание контекста для входного потока и объявление функции завершения
	outputCtx, outputCancel := context.WithCancel(context.Background()) // создание контекста для выходного потока и объявление функции завершения

	var inputWg sync.WaitGroup  // объявление WaitGroup для обеспечение выполнения горутины входного потока
	var outputWg sync.WaitGroup // объявление WaitGroup для обеспечение выполнения горутины выходного потока

	inputWg.Add(1) // увеличение счетчика WaitGroup входного потока
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done() // уменьшение счетчика WaitGroup при выходе из горутины
		for {
			// неблокирующее получение данных из канала
			select {
			case data := <-inputChannel:
				outputChannel <- data * 2 // запись данных в выходной канал
			default:
			}

			// неблокирующая проверка контекста на завершение
			select {
			case <-ctx.Done():
				fmt.Println("Input receiver stopped")
				return
			default:
			}
		}

	}(inputCtx, &inputWg)

	outputWg.Add(1) // увеличение счетчика WaitGroup
	// запуск горутины для записи данных в канал
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done() // уменьшение счетчика WaitGroup при выходе из горутины
		for {
			// неблокирующее получение данных из выходного канала
			select {
			case data := <-outputChannel:
				fmt.Println(data) // вывод данных
			default:
			}

			// неблокирующая проверка контекста на завершение
			select {
			case <-ctx.Done():
				fmt.Println("Output receiver stopped")
				return
			default:
			}
		}

	}(outputCtx, &outputWg)

	inputArray := [5]int{1, 2, 3, 4, 5} // объявление массива с исходными данными

	// запись данных во входной канал
	for _, v := range inputArray {
		inputChannel <- v
	}

	inputCancel()  // завершение контекста входного потока
	inputWg.Wait() // ожидание завершения горутины входного потока

	outputCancel()  // завершение контекста выходного потока
	outputWg.Wait() // ожидание завершения горутины выходного потока

	fmt.Println("All goroutines stopped, exiting")
}
