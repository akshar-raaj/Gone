package main

import "fmt"

func nums() (int, int) {
	/* Return two numbers, like a Python tuple*/
	x := 3
	y := 4
	return x, y
}

func main() {
	fmt.Println(nums())
}