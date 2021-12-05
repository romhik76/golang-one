package main

import (
	"fmt"
	"os"
)

func calcFactorial() {
	fmt.Println("Введите первое значение: ")
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Введен неверный операнд")
		return
	}

	fmt.Println("Введите второе значение: ")
	var b int
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Введен неверный операнд")
		return
	}

	println("Введите оператор (+, -, *, /, F(определение факториала от числа a): ")
	var operation string
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Введен неверный оператор")
		return
	}

	var result int

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
	case "F":
		result = factorial(a)
	default:
		println("некорректная операция!")
		os.Exit(1)
	}

	fmt.Printf("Результат вычисления равен: ", result)
}

func factorial(n int) int {
	sum := 1
	for i := 2; i <= n; i++ {
		sum *= i
	}
	return sum
}

func main() {
	calcFactorial()
}
