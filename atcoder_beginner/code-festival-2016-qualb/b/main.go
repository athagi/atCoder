package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	var S string
	fmt.Scan(&N, &A, &B, &S)

	counta := 0
	countb := 0
	for i := 0; i < N; i++ {
		student := S[i : i+1]
		if student == "a" {
			if counta+countb < A+B {
				fmt.Println("Yes")
				counta++
			} else {
				fmt.Println("No")
			}
		} else if student == "b" {
			if counta+countb < A+B && countb < B {
				fmt.Println("Yes")
				countb++
			} else {
				fmt.Println("No")
			}
		} else {
			fmt.Println("No")
		}
	}

}
