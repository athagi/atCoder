package main

import "fmt"

func main() {
	var A, B, X int
	fmt.Scan(&A, &B, &X)

	res := "NO"
	if A <= X && X <= A+B {
		res = "YES"
	}
	fmt.Println(res)
}
