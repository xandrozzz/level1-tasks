package tasks

import "fmt"

func GetTask13() {
	fmt.Println("Using multiple variable assignment:")
	// объявление двух целочисленных переменных
	a := 1
	b := 2
	fmt.Println("a:", a, "b:", b)
	a, b = b, a // изменение значений через объявление множественное присвоение
	fmt.Println("a:", a, "b:", b)

	fmt.Println("Using arithmetic operations:")
	// объявление двух целочисленных переменных
	c := 3
	d := 4
	fmt.Println("c:", c, "d:", d)
	// изменение значений через прибавление и вычитание чисел
	c = c + d // c=3+4=7
	d = c - d // d=7-4=3
	c = c - d // c=7-3=4
	fmt.Println("c:", c, "d:", d)

	fmt.Println("Using binary XOR:")
	// объявление двух целочисленных переменных
	e := 5
	f := 6
	fmt.Println("e:", e, "f:", f)
	// изменение значений через бинарное исключеющее или
	e = e ^ f // e=5^6=101^110=011=3
	f = e ^ f // f=3^6=011^110=101=5
	e = e ^ f // f=3^5=011^101=110=6
	fmt.Println("e:", e, "f:", f)
}
