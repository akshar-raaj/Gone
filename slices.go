package main

import "fmt"

func main() {
	arr := make([]int, 2)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	arr = append(arr, 2)
	fmt.Println(arr[2])
}