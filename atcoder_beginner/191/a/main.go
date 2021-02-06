package main

import (
	"fmt"
)

func main() {
	var V,T,S,D int
	fmt.Scan(&V, &T, &S, &D)
	res := "Yes"
	
	if T*V <= D  && D <= S*V {
		res = "No"
	}

	fmt.Println(res)


}