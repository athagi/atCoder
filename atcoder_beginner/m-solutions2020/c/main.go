package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	A := make([]int, N)
	res := ""

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		sc.Scan()
		s, _ := strconv.Atoi(sc.Text())

		A[i] = s

		if i >= K {
			if A[i-K] < A[i] {
				res = "Yes"
			} else {
				res = "No"
			}
			fmt.Println(res)
		}
	}
}
