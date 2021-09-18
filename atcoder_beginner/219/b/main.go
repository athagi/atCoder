package main

import (
	"fmt"
)

func main() {
	var S1, S2, S3 string
	var T string
	fmt.Scan(&S1, &S2, &S3, &T)

	res := ""
	for i := 0; i < len(T); i++ {
		num := T[i : i+1]
		if num == "1" {
			res += S1
		} else if num == "2" {
			res += S2
		} else {
			res += S3
		}
	}

	fmt.Println(res)

}
