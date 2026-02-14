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

	var userName string
	emptyCount := 0

	for {

		if userName == "" {
			fmt.Print("Как тебя зовут?")

			nameInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка", err)
				continue
			}
			nameInput = strings.TrimSpace(nameInput)

			if nameInput == "" {
				fmt.Println("Имя не может быть пустым")
				fmt.Println()
				continue
			}

			userName = nameInput
			fmt.Printf("Приятно познакомиться, %s!\n\n", userName)
			continue
		}

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

		//if input == "пока" {
		//	msg := fmt.Sprintf("до встречи %s, еще увидимся", userName)
		//	fmt.Println(msg)
		//	break
		//}

		input = strings.TrimSpace(input)

		if input == "сменить имя" {
			userName = ""
			fmt.Println("Ок, давай познакомимся заново 🙂")
			fmt.Println()
			continue
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
		response, shouldExit := Respond(input, userName)
		fmt.Println(response)
		fmt.Println()

		if shouldExit {
			break
		}
	}

}
