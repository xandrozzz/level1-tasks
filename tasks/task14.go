package tasks

import (
	"fmt"
	"math/rand"
	"reflect"
)

func GetTask14() {

	fmt.Println("Using reflect.TypeOf():")

	for i := 0; i < 10; i++ {
		var input interface{} // объявление переменной типа interface{}

		// присвоение переменной значение случайного типа
		x := rand.Intn(4)
		switch x {
		case 0:
			input = "aboba"
		case 1:
			input = 9412412
		case 2:
			input = true
		case 3:
			input = make(chan int)
		default:
			input = 0
		}
		fmt.Println("Value:", input, "Type:", reflect.TypeOf(input)) // вывод значения переменной и ее типа данных через функцию reflect.TypeOf()

	}

	fmt.Println("Using reflect.ValueOf.Kind():")

	for i := 0; i < 10; i++ {
		var input interface{} // объявление переменной типа interface{}

		// присвоение переменной значение случайного типа
		x := rand.Intn(4)
		switch x {
		case 0:
			input = "aboba"
		case 1:
			input = 9412412
		case 2:
			input = true
		case 3:
			input = make(chan int)
		default:
			input = 0
		}
		fmt.Println("Value:", input, "Type:", reflect.ValueOf(input).Kind()) // вывод значения переменной и ее типа данных через функцию reflect.ValueOf.Kind()

	}

	fmt.Println("Using %T format:")

	for i := 0; i < 10; i++ {
		var input interface{} // объявление переменной типа interface{}

		// присвоение переменной значение случайного типа
		x := rand.Intn(4)
		switch x {
		case 0:
			input = "aboba"
		case 1:
			input = 9412412
		case 2:
			input = true
		case 3:
			input = make(chan int)
		default:
			input = 0
		}
		fmt.Printf("Value: %v Type: %T\n", input, input) // вывод значения переменной и ее типа данных через формат %T

	}
}
