package main

import "fmt"

func swap() (int, int) {
	x := 3
	y := 4
	return x, y
}

func main() {
	fmt.Println(swap())
}