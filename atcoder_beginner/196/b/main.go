package main

import (
	"fmt"
	"strings"
)

func main() {
	var X string
	fmt.Scan(&X)

	res := strings.Split(X, ".")[0]
	fmt.Println(res)

}