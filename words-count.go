package main

import (
	"fmt"
	"strings"
)

func word_count(s string) map[string]int {
	di := make(map[string]int)
	li := strings.Fields(s)
	for _, v := range(li) {
		_, ok := di[v]
		if ok {
			di[v] = di[v] + 1
		} else {
			di[v] = 1
		}
	}
	return di
}

func main() {
	s := "my name is khan khan from epiglotis"
	result := word_count(s)
	fmt.Println(result)
}