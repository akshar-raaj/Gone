package main

import "fmt"

func change_int(x *int) {
	*x = 10
}

func change_num(x int) {
	x = 50
}

func change_array(nums [3] int) {
	/* Array is passed by value here*/
	fmt.Println("address inside function is", &nums[0])
	nums[0] = 9
}

func main() {
	var num = 100
	fmt.Println(num)

	change_int(&num)
	fmt.Println(num)

	change_num(num)
	fmt.Println(num)

	nums := [3] int{1, 2, 3}
	fmt.Println(nums)
	ptr := &nums[0]
	fmt.Println("address is", ptr)
	change_array(nums)
	fmt.Println(nums)
}