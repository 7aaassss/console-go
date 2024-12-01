package crs

import (
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Calc() {

	for {
		var a, b float64
		var op string
		fmt.Print("Введите два числа \nПервое число:")
		fmt.Scan(&a)
		fmt.Println("Введите второе число:")
		fmt.Scan(&b)
		fmt.Println("Введите операцию - (+), (-), (*), (/)")
		fmt.Scan(&op)

		switch op {
		case "+":
			fmt.Println(sum(a, b))
		case "-":
			fmt.Println(razn(a, b))
		case "*":
			fmt.Println(umn(a, b))
		case "/":
			if b == 0 {
				fmt.Println("На ноль делить нельзя!")
			} else {
				fmt.Println(del(a, b))
			}
		}
	}
}

func sum[number Number](operand1, operand2 number) number {
	return operand1 + operand2
}

func razn[number Number](operand1, operand2 number) number {
	return operand1 - operand2
}

func del[number Number](operadn1, operand2 number) number {
	return operadn1 / operand2
}

func umn[number Number](operand1, operand2 number) number {
	return operand1 * operand2
}
