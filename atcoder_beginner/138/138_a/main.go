package main

import (
	"fmt"
)

const threshold int = 3200

func main() {
	var a int
	var b string
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a >= threshold {
		fmt.Println(b)
	} else {
		fmt.Println("red")
	}
}
