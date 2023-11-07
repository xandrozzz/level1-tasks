package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func concurrentWrite(m *map[int]int, lock *sync.Mutex) {
	lock.Lock()         // вызов метода Lock для остановки других горутин
	defer lock.Unlock() // вызов метода Unlock в defer для продолжения работы других горутин после завершения текущей

	// запись данных в map
	index := rand.Intn(100)
	value := rand.Intn(100)
	fmt.Println("Writing", value, "into index", index)
	(*m)[index] = value
}

func GetTask7() {
	m := make(map[int]int) // объявление map для записи
	var lock sync.Mutex    // объявления Mutex для контроля записи в map

	ctx, cancel := context.WithCancel(context.Background()) // создание контекста и объявление функции завершения

	var wg sync.WaitGroup // объявление WaitGroup для обеспечение выполнения всех горутин

	// создание нескольких воркеров для записи в map
	for i := 0; i < 10; i++ {
		wg.Add(1) // увеличение счетчика WaitGroup
		go func(ctx context.Context, wg *sync.WaitGroup) {
			for {
				concurrentWrite(&m, &lock) // вызов функции для конкурентной записи

				time.Sleep(2 * time.Second)
				// неблокирующая проверка контекста на завершение
				select {
				case <-ctx.Done():
					fmt.Println("Writer stopped")
					wg.Done() // уменьшение счетчика WaitGroup
					return
				default:
				}
			}

		}(ctx, &wg)
	}

	time.Sleep(5 * time.Second)

	cancel() // вызов функции завершения контекста

	wg.Wait() // ожидание завершения всех горутин
	fmt.Println("All goroutines stopped")
	fmt.Println("Result map:", m)
}
