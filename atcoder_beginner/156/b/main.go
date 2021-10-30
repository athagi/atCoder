package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	fmt.Println(len(strconv.FormatInt(int64(N), K)))
}
