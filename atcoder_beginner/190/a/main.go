package main

import (
	"fmt"
)

func main() {
	var A,B,C int
	fmt.Scan(&A, &B, &C)
	res := "Takahashi"
	
	if C == 0 {
		if A - B <= 0 {
			res = "Aoki"
		}
	}
	
	if C == 1 {
		if B - A > 0 {
			res = "Aoki"
		}
	}

	fmt.Println(res)


}