package crs

import (
	"fmt"
	"math/rand"
)

func Guess() {
	fmt.Println("Давай поиграем в угадайку! Я загадал число от 1 до 100, угадай как можно быстрей")
	r := rand.Intn(100)

	count := 0
	for {
		fmt.Println("Вводи число!")
		var x int
		fmt.Scan(&x)
		count++
		if x < r {
			fmt.Println("Мое число больше!")
		} else if x > r {
			fmt.Println("Моё число меньше!")
		} else {
			fmt.Println("Поздравляю ты выиграл за ", count, " попыток!")
			break
		}

	}
}
