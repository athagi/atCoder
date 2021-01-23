package main

import (
	"fmt"
)

func main() {
	var A string
	fmt.Scan(&A)
	res := "Lost"
	s1 := A[0:1]
	s2 := A[1:2]
	s3 := A[2:3]
	if  s1 == s2 && s2 == s3 && s3 == s1{
		res = "Won"
	}
	

	fmt.Println(res)


}