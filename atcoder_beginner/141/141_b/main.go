package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a string
	fmt.Scan(&a)
	inputs := spritString(a)

	res := "Yes"
	for i, v := range inputs {
		step := i + 1
		if step%2 == 0 {
			if v == "R" {
				res = "No"
			}
		} else {
			if v == "L" {
				res = "No"
			}
		}
	}
	fmt.Println(res)
}

func spritString(input string) (res []string) {
	for i := 0; i < utf8.RuneCountInString(input); i++ {
		res = append(res, input[i:i+1])
	}
	return
}
