package main

import (
	"fmt"
)

func main() {
	var A int

	fmt.Scan(&A)
	if A < 0 {
		A = 0
	}

	fmt.Println(A)

}