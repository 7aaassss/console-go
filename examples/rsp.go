package crs

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func RSP() {
	fmt.Println("Давай сыграем в камень ножницы бумага")
	var user_choice string
	var pc_choice int // 1 - камень, 2 - бумага, 3 - ножницы
	for {
		pc_choice = rand.Intn(3)
		fmt.Println("Выбери: Камень Ножницы Бумага")
		scanner := bufio.NewScanner(os.Stdin)
		_ = scanner.Scan()
		user_choice = scanner.Text()
		choices := []string{"Камень", "Ножницы", "Бумага", "камень", "ножницы", "бумага"}
		for _, value := range choices {
			if user_choice == value {
				switch strings.ToLower(user_choice) {
				case "камень":
					if pc_choice == 1 {
						fmt.Println("Ничья!")
					} else if pc_choice == 2 {
						fmt.Println("Вы проиграли!")
					} else {
						fmt.Println("Вы выиграли!")
					}
				case "ножницы":
					if pc_choice == 3 {
						fmt.Println("Ничья!")
					} else if pc_choice == 1 {
						fmt.Println("Вы проиграли!")
					} else {
						fmt.Println("Вы выиграли!")
					}
				case "бумага":
					if pc_choice == 2 {
						fmt.Println("Ничья!")
					} else if pc_choice == 1 {
						fmt.Println("Вы проиграли!")
					} else {
						fmt.Println("Вы выиграли!")
					}
				}
			}
		}
		fmt.Println("Не валидный выбор!")
	}

}
