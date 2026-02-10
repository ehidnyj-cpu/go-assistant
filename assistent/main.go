package main

import "fmt"

func main() {
	fmt.Println("assistent:start")

	// вызов функции из logic.go
	msg := Greet("Дима")
	fmt.Println(msg)
}
