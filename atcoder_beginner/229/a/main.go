package main

import (
	"fmt"
)

func main() {
	var S1, S2 string
	fmt.Scan(&S1, &S2)

	var counter int
	if S1[0:1] == "#" {
		counter++
	}
	if S1[1:2] == "#" {
		counter++
	}
	if S2[0:1] == "#" {
		counter++
	}
	if S2[1:2] == "#" {
		counter++
	}
	res := "Yes"
	if counter >= 3 {
		res = "Yes"
	} else if (S1[0:1] == "#" && S2[1:2] == "#") || (S1[1:2] == "#" && S2[0:1] == "#") {
		res = "No"
	}
	fmt.Println(res)
}
