package tasks

import (
	"fmt"
	"math/big"
)

func GetTask8() {

	var n int64
	// получение исходного числа
	fmt.Print("Enter an int64 number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	binaryN := big.NewInt(n) // конвертация в тип big.Int

	fmt.Print("Binary representation: ")
	fmt.Printf("%b\n", binaryN)                     // вывод бинарного представления
	fmt.Println("Binary length:", binaryN.BitLen()) // получение длины бинарного представления

	var bitNum int
	// получение индекса бита для изменения
	fmt.Print("Enter the index of a bit you want to change: ")
	_, err = fmt.Scan(&bitNum)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	binaryN.SetBit(binaryN, bitNum, binaryN.Bit(bitNum)^1) // замена бита по индексу

	fmt.Print("Result binary representation: ")
	fmt.Printf("%b\n", binaryN)                     // вывод бинарного представления результата
	fmt.Println("Binary length:", binaryN.BitLen()) // получение длины бинарного представления результата

	check := binaryN.IsInt64() // проверка того, возможно ли представить результат как int64
	if !check {
		fmt.Println("The resulting number exceeds int64 boundaries:", binaryN)
	} else {
		result := binaryN.Int64() // конвертация обратно в int64
		fmt.Println("Result:", result)
	}

}
