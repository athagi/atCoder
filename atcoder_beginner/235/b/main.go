package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, h int
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	res := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		h, _ = strconv.Atoi(sc.Text())
		if res < h {
			res = h
		} else {
			break
		}
	}

	fmt.Println(res)

}
