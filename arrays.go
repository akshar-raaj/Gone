package main

import "fmt"

func main() {
	/*Define an array that can hold 5 integers*/
	var a [5]int
	a[0] = 1
	fmt.Println(a[0])
	fmt.Println(a)
	fmt.Println("Length of a is", len(a))
	fmt.Println("Capacity of a is", cap(a))
}
