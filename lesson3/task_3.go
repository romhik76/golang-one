package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	calculator()
}

func calculator() {
	fmt.Println("Введите первое значение: ")
	var a float64
	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		fmt.Println("Введен неверный операнд")
		return
	}

	fmt.Println("Введите второе значение: ")
	var b float64
	_, err = fmt.Scanf("%f", &b)
	if err != nil {
		fmt.Println("Введен неверный операнд")
		return
	}

	println("Введите оператор (+, -, *, /, ^(возведение в степень a=возводимое число,b=степень), R: ")
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
		if b == 0 {
			fmt.Println("Данная операция не может быть выполнена")
			return
		}
		result = a / b
	case "^":
		result = math.Pow(a, b)
	default:
		println("некорректная операция!")
		os.Exit(1)
	}

	fmt.Printf("Результат вычисления равен: %.2f\n", result)
}
