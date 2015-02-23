package main

import "fmt"

func main() {
	s := "akshar"
	length := len(s)
	half_length := length / 2
	rev_s := make([]rune, length)
	n := 0
	for _, r := range s {
		rev_s[n] = r
		n++
	}
	n = 0
	/*http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go#answer-1754209*/
	for n < half_length {
		rev_s[length-1-n], rev_s[n] = rev_s[n], rev_s[length-1-n]
		n++
	}
	fmt.Println(string(rev_s))
}
