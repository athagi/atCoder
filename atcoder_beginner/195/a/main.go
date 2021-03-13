package main

import (
	"fmt"
)

func main() {
	var M, H int
	fmt.Scan(&M, &H)
	res := "Yes"
	
	if H % M != 0 {
		res = "No"
	}

	fmt.Println(res)


}