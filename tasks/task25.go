package tasks

import (
	"context"
	"fmt"
	"time"
)

// собственна функция sleep, использующая контекст
func sleepWithContext(ctx context.Context, t time.Duration) {
	// ожидание завершения контекста или истечения времени
	select {
	// проверка контекста на завершение
	case <-ctx.Done():
		return
	// проверка на истечение времени
	case <-time.After(t):
		return
	}
}

func GetTask25() {

	ctx, cancel := context.WithCancel(context.Background()) //объявление контекста и функции завершения

	startTime := time.Now() // сохранение времени начала работы

	// вызов горутины, которая завершит контекст через 4 секунды
	go func(delay int, ctx context.Context, cancelFunc context.CancelFunc) {
		fmt.Println("Cancelling sleep from goroutine after", delay, "seconds")
		sleepWithContext(ctx, time.Duration(delay)*time.Second)
		cancel()
		fmt.Println("Cancel function was called")
	}(4, ctx, cancel)

	// вызов sleepWithContext на 5 секунд ожидания
	fmt.Println("Sleeping for 5 seconds")
	sleepWithContext(ctx, 5*time.Second)
	fmt.Println("Sleep finished")
	fmt.Println("Elapsed time:", time.Now().Sub(startTime)) // вывод прошедшего с начала работы времени
}
