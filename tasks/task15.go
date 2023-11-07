package tasks

import (
	"fmt"
	"math/rand"
)

var justString string // глобальная переменная для хранения строки результата

var letters = []rune("异体字異体字hmඞ?") // набор символов, включающих символы юникод для проверки

func randStringRunes(n int) string {
	b := make([]rune, n) // создание слайса рун
	// заполнение слайса рун
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b) // возврат строки, полученной из слайса рун
}

func someFunc() {
	v := []rune(randStringRunes(1 << 10)) // заполнение слайса рун

	runeArray := make([]rune, 100) // создание массива для результата

	copy(runeArray, v)             // копирование данных в новый слайс, чтобы не использовать ссылки на большую строку
	justString = string(runeArray) // конвертация слайса рун в строку

	fmt.Println("Result length:", len(runeArray), "Result string:", justString) // вывод реальной длины строки и результата

}

func GetTask15() {
	someFunc()
}
