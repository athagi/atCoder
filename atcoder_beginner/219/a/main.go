package main

import (
	"fmt"
	"strconv"
)

func main() {
	var X int
	fmt.Scan(&X)
	res := ""

	if 0 <= X && X < 40 {
		res = strconv.Itoa(40 - X)
	} else if 40 <= X && X < 70 {
		res = strconv.Itoa(70 - X)
	} else if 70 <= X && X < 90 {
		res = strconv.Itoa(90 - X)
	} else {
		res = "expert"
	}

	fmt.Println(res)

}
