package main

import (
	"fmt"
)

func main() {
	var P int
	fmt.Scan(&P)

	res := 0

	coin := make([]int, 10)
	tmp := 1
	for i := 1; i <= 10; i++ {
		tmp = tmp * i
		coin[i-1] = tmp
	}
	for i := 10; i >= 1; i-- {
		res += P / coin[i-1]
		P = P % coin[i-1]

	}
	fmt.Println(res)

}
