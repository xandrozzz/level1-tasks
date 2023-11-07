package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// структура счетчика
type safeCounter struct {
	mutex   sync.Mutex // поле mutex для ограничения доступа к счетчику из нескольких горутин одновременно
	counter int        // поле для хранения суммы счетчика
}

// Inc - метод для увеличения счетчика safeCounter
func (sc *safeCounter) Inc(v int) {
	sc.mutex.Lock()         // остановка других горутин, пытающихся вызвать метод Lock()
	defer sc.mutex.Unlock() // продолжение работы других горутин после завершения текущей

	sc.counter += v // увеличение счетчика
}

// конструктор safeCounter
func newSafeCounter() *safeCounter {
	// создание новой структуры safeCounter и возврат ссылки на нее
	return &safeCounter{
		sync.Mutex{},
		0,
	}
}

// метод для получения текущего значения счетчика
func (sc *safeCounter) getResult() int {
	return sc.counter
}

func GetTask18() {
	counter := newSafeCounter() // создание структуры safeCounter через конструктор

	var wg sync.WaitGroup // объявление WaitGroup для обеспечение выполнения всех горутин

	ctx, cancel := context.WithCancel(context.Background()) // создание контекста и объявление функции завершения

	// создание воркеров
	for i := 0; i < 10; i++ {

		wg.Add(1) // увеличение счетчика WaitGroup
		// запуск горутины воркера
		go func(counter *safeCounter, wn int, ctx context.Context, wg *sync.WaitGroup) {
			for {
				// увеличение счетчика на случайное значение
				v := rand.Intn(10)
				fmt.Println("Incrementing the counter by", v, "from worker", wn)
				counter.Inc(v)

				// неблокирующая проверка контекста на завершение
				select {
				case <-ctx.Done():
					fmt.Println("Worker", wn, "stopped")
					wg.Done() // уменьшение счетчика WaitGroup
					return
				default:
				}

				time.Sleep(time.Second)
			}

		}(counter, i, ctx, &wg)
	}

	time.Sleep(5 * time.Second)

	cancel()  // завершение контекста
	wg.Wait() // ожидание выполнения всех горутин
	fmt.Println("All goroutines stopped")

	// получение и вывод результата счетчика
	result := counter.getResult()
	fmt.Println("Counter result:", result)
}
