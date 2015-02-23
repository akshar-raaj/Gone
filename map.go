package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(m["one"])

	_, ok := m["three"]
	fmt.Println(ok)
}
