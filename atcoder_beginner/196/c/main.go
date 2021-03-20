package main

import (
		"fmt"
		"strconv"
)



func main() {
	var N int64

	fmt.Scan(&N)
	res := 0
	for i := 1; i < 1000001; i++ {
		num, _ := strconv.ParseInt(strconv.Itoa(i) + strconv.Itoa(i), 10, 64)
		if num <= N {
			res += 1
		}
	}
	fmt.Println(res)
}
