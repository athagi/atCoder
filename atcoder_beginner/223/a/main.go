package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Scan(&X)
	res := "No"

	if X%100 == 0 && X != 0 {
		res = "Yes"
	}

	fmt.Println(res)

}
