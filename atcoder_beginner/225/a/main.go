package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	a := s[0:1]
	b := s[1:2]
	c := s[2:3]

	res := 6
	if a == b && b == c && c == a {
		res = 1
	} else if a == b || b == c || c == a {
		res = 3
	}
	fmt.Println(res)
}
