package main

import (
	"fmt"
)

func main() {
	var a int
	var b string
	fmt.Scan(&a)
	fmt.Scan(&b)

	temp := ""
	res := 0

	for i := 0; i < a; i++ {
		if temp == b[i:i+1] {
			continue
		}
		temp = b[i : i+1]
		res++
	}
	fmt.Println(res)
}
