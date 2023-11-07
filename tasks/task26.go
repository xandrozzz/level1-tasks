package tasks

import (
	"fmt"
	"strings"
)

func GetTask26() {
	stringSet := make(map[rune]bool) // объявление map[rune]bool в качестве множества символов для строки

	// получение строки
	var input string
	fmt.Println("Enter a string:")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	// проверка символов на уникальность
	isUnique := true
	for _, letter := range input {
		lowercaseLetter := ([]rune(strings.ToLower(string(letter))))[0] // приведение символа к строчной форме
		// проверка наличия в множестве
		if stringSet[lowercaseLetter] {
			// если символ повторяется, установка флага isUnique на false и выход из цикла
			isUnique = false
			break
		}
		stringSet[lowercaseLetter] = true // добавление символа в множество
	}

	fmt.Println("The result is:", isUnique) // вывод результата

}
