package main

import (
	"fmt"
)

func main() {
	var digital_1, digital_2 float64
	var sign string

	fmt.Print("Введите первое число:")
	fmt.Scanln(&digital_1)

	fmt.Print("Введите второе число:")
	fmt.Scanln(&digital_2)

	fmt.Print("Введите знак для операции(+,-,/,*):")
	fmt.Scanln(&sign)

	if sign == "+" {
		meaning := digital_1 + digital_2
		fmt.Println("= ", meaning)
	} else if sign == "-" {
		meaning := digital_1 - digital_2
		fmt.Println("=", meaning)
	} else if sign == "/" {
		meaning := digital_1 / digital_2
		fmt.Println("=", meaning)
	} else if sign == "*" {
		meaning := digital_1 * digital_2
		fmt.Println("=", meaning)
	}
}
