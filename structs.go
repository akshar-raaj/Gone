package main

import "fmt"

type Vertex struct {
	x int
	y int
	z *int
}

func main() {
	num := 5
	v := Vertex{x: 1, y: 2, z: &num}
	fmt.Println(v)
	v.x = 4
	fmt.Println(v)
	ptr := &v
	fmt.Println(*ptr)

	fmt.Println("********************")
	num2 := 6
	v2 := Vertex{4, 5, &num2}
	*ptr = v2
	fmt.Println(v2)
	fmt.Println(v)
}
