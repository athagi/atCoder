package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)
	a := 60*60*A + 60*B
	b := 60*60*C + 60*D + 1
	res := "Takahashi"
	if a > b {
		res = "Aoki"
	}
	fmt.Println(res)
}
