package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("Word", "Another word")
	fmt.Println(swap(1, 2))
	a, b := swap(4, 5)
	fmt.Println(a, b)
}