package crs

import (
	"encoding/json"
	"log"
	"os"
)

type Pipen struct { // структура описывающая жсон и работу с данными пользователя
	ABC string
}

func LOL() { // главная экспортируемая функция

}
func upper(mainList *[]Pipen) {
	content, _ := json.Marshal(mainList)
	err := os.WriteFile("etc.json", content, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func getdata() []Pipen { // функция для работы с жсон - возвращает список пользователей
	var listy []Pipen
	content, _ := os.ReadFile("etc.json")
	body := json.Unmarshal(content, &listy)
	if body != nil {
		log.Fatal("Error during Unmarshal(): ", body)
	}
	return listy
}
