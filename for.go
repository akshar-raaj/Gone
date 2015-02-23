package main

import "fmt"

func main() {
	i := 0
	for i<10 {
		fmt.Println(i)
		i += 1
	}

	arr := []int{3, 6, 9}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}