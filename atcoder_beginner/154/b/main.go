package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	res := ""
	for i := 0; i < len(S); i++ {
		res += "x"
	}
	fmt.Println(res)
}
