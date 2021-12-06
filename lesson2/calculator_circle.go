package main

import (
	"fmt"
	"math"
)

func circle() {

	var s float64
	fmt.Println("Введите площадь окружности: ")
	fmt.Scan(&s)
	d := 2 * math.Sqrt(s/math.Pi)
	с := d * math.Pi
	fmt.Println(" диаметр окружности равен: ", d, "\n", "длинна окружности равна: ", с)
}

func main() {
	circle()
}
