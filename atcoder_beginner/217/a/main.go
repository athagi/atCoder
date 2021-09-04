package main

import (
	"fmt"
)

func main() {
	var S, T string
	fmt.Scan(&S, &T)
	res := "No"

	if S < T {
		res = "Yes"
	}

	fmt.Println(res)

}
