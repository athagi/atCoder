package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	res := "Yes"
	if len(S)%2 == 0 {
		for i := 0; i < len(S); i += 2 {
			char := S[i : i+2]
			if char != "hi" {
				res = "No"
			}
		}
	} else {
		res = "No"
	}
	fmt.Println(res)
}
