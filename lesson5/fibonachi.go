package main

import "fmt"

func fibonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonachi(n-1) + fibonachi(n-2)
}

func number() {
	var n uint
	fmt.Println("Введите число n")
	fmt.Scanln(&n)

	fmt.Println(fmt.Sprintf("result for %d: %d", n, fibonachi(n)))
	number()
}
func main() {
	number()

}
