package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A, B int64
	fmt.Scan(&A, &B)
	a := fmt.Sprintf("%018d", A)
	b := fmt.Sprintf("%018d", B)

	res := "Easy"
	for i := 0; i < 18; i++ {
		x, _ := strconv.Atoi(a[i : i+1])
		y, _ := strconv.Atoi(b[i : i+1])
		if x+y >= 10 {
			res = "Hard"
		}
	}

	fmt.Println(res)

}
