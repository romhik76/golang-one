package main

import "fmt"

func main() {
	fmt.Println("Введите значение: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Введено неверное значение")
		return
	}
	for j := 2; j < n; j++ {
		var flag = 0
		for i := 2; i <= j/2; i++ {
			if j%i == 0 {
				flag = 1
				break
			}

		}
		if flag == 0 {
			fmt.Println(j)
		}
	}
}
