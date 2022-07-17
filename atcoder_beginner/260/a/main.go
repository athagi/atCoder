package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	a := S[0:1]
	b := S[1:2]
	c := S[2:3]

	res := a
	if a == b && b == c && c == a {
		res = "-1"
	} else if a == b {
		res = c
	} else if b == c {
		res = a
	} else if c == a {
		res = b
	}

	fmt.Println(res)
}
