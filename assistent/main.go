package main

import (
	"bufio" //пакет для чтения ввода
	"fmt"
	"os"      // доступ к стандартному вводу (stdin)
	"strings" // работа со строками
	//для подсчета символов на кириллице
)

func main() {

	fmt.Println()
	fmt.Println("assistent:start")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin) // Создаем reader для чтения файла

	for {

		fmt.Print("Введи команду: ")

		input, err := reader.ReadString('\n')
		fmt.Println()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "выход" {
			fmt.Println("Завершение работы")
			break
		}

		response := Respond(input)
		fmt.Println(response)
		fmt.Println()
	}

}
