package main

import (
	"fmt"
)

func main() {
	var s string

	S := make(map[string]bool)
	for i := 0; i < 4; i++ {
		fmt.Scan(&s)
		if !S[s] {
			S[s] = true
		}
	}
	res := "No"
	if len(S) == 4 {
		res = "Yes"
	}
	fmt.Println(res)

}
