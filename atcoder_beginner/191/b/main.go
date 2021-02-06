package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var N, X, A int
	var tmp string
	fmt.Scan(&N, &X)
	res := []string{}
	for i:= 0; i < N; i++{
		fmt.Scan(&A)
		if A != X {
			tmp = strconv.Itoa(A)
			res = append(res, tmp)
		}
	}

	fmt.Println(strings.Join(res, " "))


}