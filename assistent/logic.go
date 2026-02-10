package main

import "fmt"

func Greet(name string) string {
	u := User{Name: name} //struct is model.go
	return fmt.Sprintf("Привет, %s! Проект живой ✅", u.Name)
}
