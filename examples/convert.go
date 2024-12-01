package crs

import "fmt"

func Convert() {
	for {
		fmt.Println("Выберите что конвертировать?")
		fmt.Println("1 - Градусы, 2 -Длина, 3 - Время, 4 - Масса")
		var what int

		fmt.Scan(&what)

		switch what {
		case 1:
			fmt.Println("Градусы в кельвины - 1, Градусы в Фаренгейт - 2")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				fmt.Println("Введите градусы")
				var gradus float64
				fmt.Scan(&gradus)
				fmt.Println(gradus, " градусов цельсия равно ", gradus+273, " кельвинов")
			case 2:
				fmt.Println("Введите градусы")
				var gradus float64
				fmt.Scan(&gradus)
				fmt.Println(gradus, " градусов цельсия равно ", (32.0 + 9.0*gradus/5.0), " градусов по фаренгейту")
			}
		case 2:
			fmt.Println("КМ в метры - 1, Метры в КМ - 2")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				fmt.Println("Введите км")
				var km float64
				fmt.Scan(&km)
				fmt.Println(km, " километров равно ", km*1000.0)
			case 2:
				fmt.Println("Введите метры")
				var meters float64
				fmt.Scan(&meters)
				fmt.Println(meters, " метров равно ", meters/1000.0)
			}
		case 3:
			fmt.Println("Часы в минуты - 1, часы в секунды - 2, минуты в секунды - 3")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				fmt.Println("Введите часы")
				var hours float64
				fmt.Scan(&hours)
				fmt.Println(hours, " часов равно ", hours*60.0, " минут")
			case 2:
				fmt.Println("Введите часы")
				var hours float64
				fmt.Scan(&hours)
				fmt.Println(hours, " часов равно ", hours*3600.0, " секунд")
			case 3:
				fmt.Println("Введите минуты")
				var minutes float64
				fmt.Scan(&minutes)
				fmt.Println(minutes, " минут равно ", minutes*60.0, " секунд")
			}
		case 4:
			fmt.Println("КГ в граммы - 1, граммы в КГ - 2")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				fmt.Println("Введите КГ")
				var kg float64
				fmt.Scan(&kg)
				fmt.Println(kg, " килограммов равно ", kg*1000.0, " грамм")
			case 2:
				fmt.Println("Введите граммы")
				var gm float64
				fmt.Scan(&gm)
				fmt.Println(gm, " граммов равно ", gm/1000.0, " килограмм")
			}
		}
	}
}
