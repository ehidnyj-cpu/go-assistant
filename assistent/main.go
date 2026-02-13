package main

import (
	"bufio" //пакет для чтения ввода
	"fmt"
	"os"      // доступ к стандартному вводу (stdin)
	"strings" // работа со строками
)

func main() {

	fmt.Println()
	fmt.Println("assistent:start")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin) // Создаем reader для чтения файла

	emptyCount := 0

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

		input = strings.TrimSpace(input)

		if input == "" {
			emptyCount++

			if emptyCount == 1 {
				fmt.Println("Ну давай пообщаемся")
				fmt.Println()
			}

			if emptyCount == 2 {
				fmt.Println("Эй, ты здесь?")
				fmt.Println()
			}

			if emptyCount == 3 {
				fmt.Println("Ну не хочешь общаться и ладно. Пока")
				break
			}

			continue
		}
		emptyCount = 0
		response := Respond(input)
		fmt.Println(response)
		fmt.Println()
	}

}
