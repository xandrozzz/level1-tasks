package tasks

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

var letterRunesSpace = []rune("异体字異体字hmඞ?abo ") // набор символов, включающих символы юникод для проверки, а также пробел

func randStringWordsUnicode(n int) string {
	b := make([]rune, n) // создание слайса рун
	// заполнение слайса рун
	for i := range b {
		b[i] = letterRunesSpace[rand.Intn(len(letterRunesSpace))]
	}
	return string(b) // возврат строки, полученной из слайса рун
}

func GetTask20() {
	input := randStringWordsUnicode(30) // генерация случайной строки с символами юникод и пробелами

	fmt.Println("Using slices.Reverse():")

	inputWords := strings.Split(input, " ") // разделение строки на слайс строк (слов)
	// вывод изначальной строки
	fmt.Println("Initial string:")
	fmt.Println(input)

	slices.Reverse(inputWords) // переворот слайса
	// вывод развернутого слайса, соединенного обратно в строку
	fmt.Println("Reversed string:")
	fmt.Println(strings.Join(inputWords, " "))

	fmt.Println("Using builtin methods:")

	inputWords = strings.Split(input, " ") // разделение строки на слайс строк (слов)
	// вывод изначальной строки
	fmt.Println("Initial string:")
	fmt.Println(input)

	length := len(inputWords) // получение длины слайса
	// цикл для переворота слайса
	for i := 0; i < (length)/2; i++ {
		inputWords[i], inputWords[length-i-1] = inputWords[length-i-1], inputWords[i] // обмен местами элементов на противоположных концах слайса
	}

	// вывод развернутого слайса, соединенного обратно в строку
	fmt.Println("Reversed string:")
	fmt.Println(strings.Join(inputWords, " "))
}
