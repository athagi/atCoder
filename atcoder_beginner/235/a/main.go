package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)
	a := S[0:1]
	b := S[1:2]
	c := S[2:3]
	x, _ := strconv.Atoi(a + b + c)
	y, _ := strconv.Atoi(b + c + a)
	z, _ := strconv.Atoi(c + a + b)

	fmt.Println(x + y + z)
}
