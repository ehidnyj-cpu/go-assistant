package main

import (
	"fmt"
	"strings"
	"time"
)

func Respond(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))

	if strings.Contains(input, "привет") {
		return "Привет, рад тебя видеть"
	}

	if strings.Contains(input, "время") || strings.Contains(input, "час") {
		return fmt.Sprintf("Сейчас %s", time.Now().Format("15:04:05"))
	}

	if strings.Contains(input, "пока") {
		return "до встречи"
	}

	if strings.Contains(input, "кто ты") {
		return "Я твой Go ассистент"
	}

	if strings.Contains(input, "как дела") {
		return "Отлично!"
	}

	return "Я пока не понимаю эту команду"

}
