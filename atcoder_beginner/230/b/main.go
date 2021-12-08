package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	T := "oxxoxxoxxoxxoxxoxxoxxoxxoxxoxx"
	res := "No"
	if strings.Contains(T, S) {
		res = "Yes"
	}
	fmt.Println(res)
}
