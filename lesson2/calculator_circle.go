package main

import (
	"fmt"
)

func circle() {
	const pi = 3.14

	var s float32
	fmt.Println("Введите площадь окружности: ")
	fmt.Scan(&s)
	a := (s / pi) / (s / pi)
	d := 2 * a
	с := d * pi
	fmt.Println(" диамитер окружности равен: ", d, "\n", "длинна окружности равна: ", с)
}

func main() {
	circle()
}
