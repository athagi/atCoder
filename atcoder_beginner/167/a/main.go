package main

import (
	"fmt"
)

func main() {
	var S, T string
	fmt.Scan(&S, &T)

	res := "No"
	if (len(S) + 1 == len(T)) && T[0:len(T)-1] == S {
		res = "Yes"
	}
	

	fmt.Println(res)
}
