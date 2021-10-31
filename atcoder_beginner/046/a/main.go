package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	p := []int{a, b, c}
	res := make(map[int]bool)
	for i := 0; i < len(p); i++ {
		if !res[p[i]] {
			res[p[i]] = true
		}
	}

	fmt.Println(len(res))
}
