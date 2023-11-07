package tasks

import (
	"fmt"
	"math"
)

// структура точки
type point struct {
	// поля с координатами точки
	x float64
	y float64
}

// конструктор точки
func newPoint(x, y float64) *point {
	return &point{x, y} // возврат указателя на объект point с координатами, полученными из параметров
}

// метод для нахождения расстояния между двумя точками
func (p *point) findDistance(anotherPoint *point) float64 {
	return math.Sqrt(math.Pow(p.x-anotherPoint.x, 2) + math.Pow(p.y-anotherPoint.y, 2)) // возврат расстояния между двумя точками, полученного по формуле
}

func GetTask24() {
	// получение координат первой точки
	var x1, y1 float64
	fmt.Println("Enter the coordinates of the first point:")
	_, err := fmt.Scan(&x1, &y1)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	// получение координат второй точки
	var x2, y2 float64
	fmt.Println("Enter the coordinates of the second point:")
	_, err = fmt.Scan(&x2, &y2)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	// объявление двух объектов point через конструктор
	point1 := newPoint(x1, y1)
	point2 := newPoint(x2, y2)

	// вывод расстояния между точками
	fmt.Println("The distance between two points is:", point1.findDistance(point2))
}
