package crs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var filename string = "users.json"

type User struct { // структура описывающая жсон и работу с данными пользователя
	Login    string
	Password string
}

type UsersManager struct {
	Users []User
}

func (manage *UsersManager) AddNew(login, password string) {
	newUser := User{
		Login:    login,
		Password: password,
	}
	manage.Users = append(manage.Users, newUser)
}

func reg(manager *UsersManager) int {
	fmt.Println("Чтобы добавить пользователя введите логин, а затем пароль")
	_ = scanner.Scan()
	login := scanner.Text()
	cnt := 0
	for i := 0; i < len(manager.Users); i++ {
		if manager.Users[i].Login == login {
			fmt.Println("Пользователь с таким логином уже существует!, Введите другой")
			_ = scanner.Scan()
			login = scanner.Text()
			cnt++
			if cnt == 3 {
				return 0
			}
		}
	}
	_ = scanner.Scan()
	password := scanner.Text()

	manager.AddNew(login, password)
	newJson(manager)
	return 1
}

func App() { // главная экспортируемая функция
	for {
		fmt.Println("Добро пожаловать в систему управления пользователями. Для начала выберите: авторизация - 1, регистрация - 2. Если вы хотите выйте напишите q")
		mainList, err := up() // инит главного слайса структур с которым работаем
		manager := UsersManager{Users: mainList}
		if err != nil {
			fmt.Print(err)
		}
		_ = scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			fmt.Println("Введите пароль и логин")
			_ = scanner.Scan()
			login := scanner.Text()
			_ = scanner.Scan()
			password := scanner.Text()
			x, y := checkUser(&login, &password, &mainList)
			if x && y {
				fmt.Println("Вы успешно авторизовались!")
			} else if x {
				fmt.Println("Вы ввели неверный пароль!")
			} else {
				fmt.Println("Такой пользователь не найден")
			}
		case "2":
			xd := reg(&manager)
			if xd == 0 {
				fmt.Println("Вы исчерпали все свои попытки")
			}
		case "q":
			os.Exit(1)
		}
	}
}

func up() ([]User, error) { // функция для работы с жсон - возвращает список пользователей
	var listy []User
	content, err := os.ReadFile(filename)
	if err != nil {
		return listy, err
	}

	jerr := json.Unmarshal(content, &listy)
	if jerr != nil {
		return listy, jerr
	}
	return listy, nil
}

func checkUser(login, password *string, listy *[]User) (bool, bool) { // проверяем данные пользователя; ретерним булеан
	for _, value := range *listy {
		if value.Login == *login {
			if value.Password == *password {
				return true, true
			} else {
				return true, false
			}
		}
	}
	return false, false
}

func newJson(manager *UsersManager) {
	content, _ := json.Marshal(manager.Users)
	err := os.WriteFile(filename, content, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
