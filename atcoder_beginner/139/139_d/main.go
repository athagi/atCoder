package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	res := a * (a - 1) / 2
	fmt.Println(res)

}
