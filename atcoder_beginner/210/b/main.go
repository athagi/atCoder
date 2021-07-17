package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	res := "Takahashi"
	for i := 0; i < N; i++ {
		s := S[i : i+1]
		if s == "1" {
			if i%2 == 1 {
				res = "Aoki"
				break
			} else {
				break
			}
		}
	}

	fmt.Println(res)

}
