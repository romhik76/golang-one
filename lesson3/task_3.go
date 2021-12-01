package main

import (
	"fmt"
	"os"
)

func calculator() {
	fmt.Println("Введите первое число: ")
	var a float64
	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		fmt.Println("Введен неверный операнд")
		return
	}

	fmt.Println("Введите второе число: ")
	var b float64
	_, err = fmt.Scanf("%f", &b)
	if err != nil {
		fmt.Println("Введен неверный операнд")
		return
	}

	println("Введите оператор (+, -, *, /): ")
	var operation string
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Введен неверный оператор")
		return
	}

	var result float64

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
		if a == 0 || b == 0 {
			fmt.Println("Данная операция не может быть выполнена")
			return
		}
	default:
		println("некорректная операция!")
		os.Exit(1)
	}

	fmt.Println("Результат вычисления равен:", result)
}
func main() {
	calculator()
}
