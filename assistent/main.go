package main

import (
	"bufio" //пакет для чтения ввода
	"fmt"
	"os"           // доступ к стандартному вводу (stdin)
	"strings"      // работа со строками
	"unicode/utf8" //для подсчета символов на кириллице
)

func main() {

	fmt.Println()
	fmt.Println("assistent:start")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin) // Создаем reader для чтения файла

	var name string

	for {

		fmt.Println("Как тебя зовут?") // Print без перевода строки
		fmt.Println()
		name, _ = reader.ReadString('\n') // Читаем строку до нажатия Enter

		name = strings.TrimSpace(name) //убираем \n и лишние пробелы

		fmt.Println()

		if name == "" {
			fmt.Println("Имя не введено, попробуй ещё раз.")
			fmt.Println()
			continue
		}

		length := utf8.RuneCountInString(name)

		if length < 2 {
			fmt.Println("Слишком коротко для имени :)")
			fmt.Println()
			continue
		}

		if length > 15 {
			fmt.Println("Слишком длинное имя я так не запомню :)")
			fmt.Println()
			continue
		}

		msg := Greet(name) // Вызываем нашу функцию логики
		fmt.Println(msg)
		fmt.Println()
		break
	}

	fmt.Println("len:", len(name))
	fmt.Println("runes", utf8.RuneCountInString(name))

}
