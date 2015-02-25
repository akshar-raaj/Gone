package main

import "fmt"

func change_slice(slice []int) {
	//Slice only points to internal array.
	//So the passed array will be changed and change will
	//be reflected in the passed slice
	slice[0] = 50
}

func main() {
	arr := make([]int, 2)
	arr[0] = 5
	arr[1] = 3
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	arr = append(arr, 2)

	change_slice(arr)
	fmt.Println("*******")
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	new_arr := arr[1:3]
	
	change_slice(new_arr)
	fmt.Println("*******")
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
}
