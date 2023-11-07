package tasks

import (
	"fmt"
	"math/big"
)

func GetTask22() {
	// объявление переменных для хранения переменных a и b, а также результатов операций
	var a, b, c big.Int
	var d big.Float

	// получение числа a
	fmt.Print("Enter a: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	// получение числа b
	fmt.Print("Enter b: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	fmt.Println("a + b =", c.Add(&a, &b))                                                 // вычисление суммы чисел a и b
	fmt.Println("a - b =", c.Sub(&a, &b))                                                 // вычисление разности чисел a и b
	fmt.Println("a * b =", c.Mul(&a, &b))                                                 // вычисление произведения чисел a и b
	fmt.Println("a / b =", d.Quo(big.NewFloat(0).SetInt(&a), big.NewFloat(0).SetInt(&b))) // конвертация чисел a и b в big.Float и нахождение их частного
}
