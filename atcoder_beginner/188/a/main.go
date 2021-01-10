package main

import (
	"fmt"
)

func main() {
	var A,B int
	fmt.Scan(&A,&B)
	res := "No"
	if A - B < 3 && A - B > -3 {
		res = "Yes"
	}
	

	fmt.Println(res)


}