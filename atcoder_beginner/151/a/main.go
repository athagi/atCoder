package main

import (
	"fmt"
)

func main() {
	var c string
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	fmt.Scan(&c)
	res := "a"
	for i := 0; i < len(chars); i++ {
		if chars[i] == c {
			res = chars[i+1]
			break
		}
	}

	fmt.Println(res)
}
