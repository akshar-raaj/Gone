package main

import "fmt"

func main() {
	s := "akshar"
	fmt.Println(s)
	fmt.Println(len(s))
	rev_s := make([]byte, len(s))
	for _, v := range(s) {
		fmt.Println(v)
		rev_s.append(v)
	}
}