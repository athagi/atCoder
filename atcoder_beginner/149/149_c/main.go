package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Scan(&X)

	x := X
	res := 2
	if x != 2 {
		if X%2 == 0 {
			x = X + 1
		}
		for i := x; ; i += 2 {
			flag := true
			for j := 3; j < X; j += 2 {
				if i%j == 0 {
					flag = false
					break
				}
			}
			if flag {
				res = i
				break
			}
		}
	}
	fmt.Println(res)
}
