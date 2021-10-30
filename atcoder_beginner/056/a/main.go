package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	res := ""
	if a == "H" {
		if b == "H" {
			res = "H"
		} else {
			res = "D"
		}
	} else {
		if b == "H" {
			res = "D"
		} else {
			res = "H"
		}
	}
	fmt.Println(res)
}
