package main

import "fmt"

func main() {
	var R, L int
	str := "atcoder"
	fmt.Scan(&R, &L)
	fmt.Println(str[R-1 : L])
}
