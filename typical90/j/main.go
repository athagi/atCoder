package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, C, P, Q, L, R int
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sum1 := make([]int, N+1)
	sum1[0] = 0
	sum2 := make([]int, N+1)
	sum2[0] = 0
	for i := 1; i < N+1; i++ {
		sc.Scan()
		C, _ = strconv.Atoi(sc.Text())

		sc.Scan()
		P, _ = strconv.Atoi(sc.Text())

		if C == 1 {
			sum1[i] = sum1[i-1] + P
			sum2[i] = sum2[i-1]
		} else {
			sum2[i] = sum2[i-1] + P
			sum1[i] = sum1[i-1]
		}
	}

	sc.Scan()
	Q, _ = strconv.Atoi(sc.Text())
	for i := 0; i < Q; i++ {
		sc.Scan()
		L, _ = strconv.Atoi(sc.Text())

		sc.Scan()
		R, _ = strconv.Atoi(sc.Text())
		class1 := sum1[R] - sum1[L-1]
		class2 := sum2[R] - sum2[L-1]
		fmt.Println(class1, class2)
	}
}
