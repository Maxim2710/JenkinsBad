package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем срез с разными приветствиями
	greetings := []string{
		"Привет, мир!",
		"Приветствую тебя!",
		"ДАРОВА!",
		"Привет, Дима!",
	}

	// Выбираем случайное приветствие из среза
	randomIndex := rand.Intn(len(greetings))
	randomGreeting := greetings[randomIndex]

	// Выводим случайное приветствие
	fmt.Println(randomGreeting)
}

