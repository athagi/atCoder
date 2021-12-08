package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	res := N
	if N >= 42 {
		res = N + 1
	}
	s := fmt.Sprintf("%03d", res)
	fmt.Println("AGC" + s)
}
