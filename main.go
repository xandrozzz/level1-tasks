package main

import (
	"fmt"
	"level1-tasks/tasks"
)

func main() {

	var i int

	fmt.Print("Enter a task number: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}
	tasks.GetTask(i)
}
