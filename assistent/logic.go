package main

import (
	"fmt"
	"strings"
	"time"
)

func Respond(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))

	switch input {

	case "привет":
		return "Привет! Рад тебя видеть!"

	case "пока":
		return "До встречи"

	case "время":
		return fmt.Sprintf("Сейчас %s", time.Now().Format("15:04:05"))

	case "hola":
		return "Buenos"

	case "кто ты":
		return "Я твой ассистент"

	default:
		return "Я пока не понимаю эту команду"

	}

}
