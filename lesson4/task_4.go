package main

import "fmt"

func main() {

	var n int
	fmt.Println("Введите размер массива:")
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		fmt.Println(fmt.Sprintf("Err: %s", err.Error()))
	}

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Ведите число: ")
		_, err := fmt.Scanf("%d", &slice[i])
		if err != nil {
			fmt.Println(fmt.Sprintf("Err: %s", err.Error()))
		}
	}
	fmt.Println(fmt.Sprintln("Введены:", slice))

	for i := 1; i < len(slice); i++ {
		current := slice[i]
		j := i
		for j > 0 && slice[j-1] > current {
			slice[j] = slice[j-1]
			j--
		}
		slice[j] = current
	}

	fmt.Println("Сортировка:", slice)
}
