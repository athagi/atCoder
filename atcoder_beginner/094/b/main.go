package main

import "fmt"

func main() {
	var A, B, X int
	fmt.Scan(&A, &B, &X)

	res := "No"
	if A <= X && X <= A+B {
		res = "Yes"
	}
	fmt.Println(res)
}
