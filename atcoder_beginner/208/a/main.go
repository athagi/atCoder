package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	res := "No"
	if A <= B && A*6 >= B {
		res = "Yes"
	}

	fmt.Println(res)

}
