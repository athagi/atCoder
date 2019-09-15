package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)

	weathers := []string{"Sunny", "Cloudy", "Rainy"}

	num := 0
	if a == "Sunny" {
		num = 0
	} else if a == "Cloudy" {
		num = 1
	} else {
		num = 2
	}
	res := num + 1
	if res == 3 {
		res = 0
	}

	fmt.Println(weathers[res])
}
