package tasks

import "fmt"

// создание структуры Human

type Human struct {
	Age  int
	Name string
}

// создание структуры Action

type Action struct {
	Human //встраивание структуры Human в структуру Action
}

//объявление методов для структуры Human

func (h *Human) SayAge() {
	fmt.Println("I'm", h.Age, "years old")
}

func (h *Human) SayName() {
	fmt.Println("My name is", h.Name)
}

func (h *Human) Sleep() {
	fmt.Println("Zzzzzzzz...")
}

func GetTask1() {
	//создание экземпляров структур

	human := Human{Age: 20, Name: "Igor"}
	action := Human{38, "Vitya"}

	fmt.Println("Human:")

	//вызов методов структуры Human

	human.SayAge()
	human.SayName()
	human.Sleep()

	fmt.Println("Action:")

	//вызов методов структуры Action

	action.SayAge()
	action.SayName()
	action.Sleep()
}
