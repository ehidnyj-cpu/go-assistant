package main

import (
	"fmt"
	"strings"
	"time"
)

func Respond(input string, userName string) (string, bool) {
	input = strings.ToLower(strings.TrimSpace(input))

	if strings.Contains(input, "привет") {
		return fmt.Sprintf("Привет, %s! рад тебя видеть", userName), false
	}

	if strings.Contains(input, "время") || strings.Contains(input, "час") {
		return fmt.Sprintf("Сейчас %s", time.Now().Format("15:04:05")), false
	}

	if strings.Contains(input, "пока") {
		return fmt.Sprintf("до встречи %s, еще увидимся", userName), true
	}

	if strings.Contains(input, "кто ты") {
		return fmt.Sprintf("Я твой друг, %s ", userName), false
	}

	if strings.Contains(input, "как дела") {
		return fmt.Sprintf("Отлично, %s", userName), false
	}

	return "userName, Я пока не понимаю эту команду", false

}
