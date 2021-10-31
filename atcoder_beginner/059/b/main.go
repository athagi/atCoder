package main

import "fmt"

func main() {
	var A, B string
	fmt.Scan(&A, &B)
	res := ""
	if len(A) > len(B) {
		res = "GREATER"
	} else if len(A) < len(B) {
		res = "LESS"
	} else {
		if A > B {
			res = "GREATER"
		} else if A < B {
			res = "LESS"
		} else {
			res = "EQUAL"
		}
	}

	fmt.Println(res)
}
