package tasks

import (
	"fmt"
	"math/rand"
	"slices"
)

var letterRunes = []rune("异体字異体字hmඞ?") // набор символов, включающих символы юникод для проверки

func randStringRunesUnicode(n int) string {
	b := make([]rune, n) // создание слайса рун
	// заполнение слайса рун
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b) // возврат строки, полученной из слайса рун
}

func GetTask19() {
	input := randStringRunesUnicode(20) // генерация случайной строки с символами юникод

	fmt.Println("Using slices.Reverse():")

	inputRunes := []rune(input) // конвертация строки в слайс рун
	// вывод изначального слайса, конвертированного в строку
	fmt.Println("Initial string:")
	fmt.Println(string(inputRunes))

	slices.Reverse(inputRunes) // переворот слайса
	// вывод развернутого слайса, сконвертированного в строку
	fmt.Println("Reversed string:")
	fmt.Println(string(inputRunes))

	fmt.Println("Using builtin methods:")

	inputRunes = []rune(input) // конвертация строки в слайс рун
	// вывод изначального слайса, конвертированного в строку
	fmt.Println("Initial string:")
	fmt.Println(string(inputRunes))

	length := len(inputRunes) // получение длины слайса
	// цикл для переворота слайса
	for i := 0; i < (length)/2; i++ {
		inputRunes[i], inputRunes[length-i-1] = inputRunes[length-i-1], inputRunes[i] // обмен местами элементов на противоположных концах слайса
	}

	// вывод развернутого слайса, сконвертированного в строку
	fmt.Println("Reversed string:")
	fmt.Println(string(inputRunes))

}
