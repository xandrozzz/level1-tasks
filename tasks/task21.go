package tasks

import (
	"fmt"
	"reflect"
)

// IntPrinter - класс принтера чисел
type IntPrinter struct{}

// PrintInt - метод вывода числа
func (pr *IntPrinter) PrintInt(number int) {
	fmt.Println(number) // вывод числа
}

// StringPrinter - класс принтера строк
type StringPrinter struct{}

// PrintString - метод вывода строки
func (pr *StringPrinter) PrintString(line string) {
	fmt.Println(line) // вывод строки
}

// PrinterAdapter - целевой интерфейс, который должны реализовать адаптеры
type PrinterAdapter interface {
	PrintWhatever() // объявление метода вывода для адаптера
}

// IntPrinterAdapter - адаптер для IntPrinter
type IntPrinterAdapter struct {
	*IntPrinter // анонимное поле со ссылкой на изначальный класс
}

// PrintWhatever - реализация метода PrintWhatever у класса IntPrinterAdapter
func (adapter *IntPrinterAdapter) PrintWhatever() {
	adapter.PrintInt(322) // вызов метода класса IntPrinter
}

// NewIntPrinterAdapter - конструктор адаптера IntPrinterAdapter
func NewIntPrinterAdapter(pr *IntPrinter) PrinterAdapter {
	return &IntPrinterAdapter{pr} // возвращение ссылки на объект класса IntPrinterAdapter, реализующий интерфейс PrinterAdapter
}

// StringPrinterAdapter - адаптер для StringPrinter
type StringPrinterAdapter struct {
	*StringPrinter // анонимное поле со ссылкой на изначальный класс
}

// PrintWhatever - реализация метода PrintWhatever у класса StringPrinterAdapter
func (adapter *StringPrinterAdapter) PrintWhatever() {
	adapter.PrintString("aboba") // вызов метода класса StringPrinter
}

// NewStringPrinterAdapter - конструктор адаптера StringPrinterAdapter
func NewStringPrinterAdapter(pr *StringPrinter) PrinterAdapter {
	return &StringPrinterAdapter{pr} // возвращение ссылки на объект класса StringPrinterAdapter, реализующий интерфейс PrinterAdapter
}

// AmogusPrinter - класс принтера символов ඞ
type AmogusPrinter struct{}

// PrintWhatever - метод вывода символа ඞ
func (w *AmogusPrinter) PrintWhatever() {
	fmt.Println("ඞ") // вывод символа ඞ
}

func GetTask21() {
	// объявление слайса, состоящего из объектов IntPrinterAdapter, StringPrinterAdapter и AmogusPrinter, каждый из которых реализует интерфейс PrinterAdapter
	printerSlice := []PrinterAdapter{NewIntPrinterAdapter(&IntPrinter{}), NewStringPrinterAdapter(&StringPrinter{}), &AmogusPrinter{}}

	// цикл для поочередного вызова метода PrintWhatever() у каждого из объектов
	for _, printer := range printerSlice {
		fmt.Println("Class:", reflect.TypeOf(printer)) // вывод имени класса объекта
		printer.PrintWhatever()                        // вызов метода PrintWhatever()
	}
}
