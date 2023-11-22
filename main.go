package main

import (
	"fmt"
)

func main() {
	print("Вызов")
	print("asda")
	print("safew")
}

func print(message string) {
	fmt.Println(sayHello(message, 12))
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет %s! Тебе %d лет!", name, age)
}