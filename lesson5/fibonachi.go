package main

import "fmt"

var myMap = make(map[uint]uint)

func Fibonachi(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	val, ok := myMap[n]
	if ok {
		fmt.Println(fmt.Sprintf("map: %d", myMap[n]))

		return val
	}
	myMap[n] = Fibonachi(n-1) + Fibonachi(n-2)
	fmt.Println(fmt.Sprintf("fin: %d", myMap[n]))

	return myMap[n]
}

func Number() {
	var n uint
	fmt.Println("Введите число n")
	fmt.Scanln(&n)

	fmt.Println(fmt.Sprintf("result for %d: %d", n, Fibonachi(n)))
	Number()
}
func main() {
	Number()

}
