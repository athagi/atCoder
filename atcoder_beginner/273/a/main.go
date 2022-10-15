package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(f(N))
}

func f(k int) int {
	if k == 0 {
		return 1
	}
	return k * f(k-1)
}
