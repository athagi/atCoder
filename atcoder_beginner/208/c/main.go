package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var N int
	var K int64
	fmt.Scan(&N, &K)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		A[i] = a
	}

	min := K / int64(N)
	mod := K % int64(N)

	B := make([]int, N)
	copy(B, A)
	sort.Ints(B)

	num := int(mod)
	if mod > 0 {
		num = B[mod-1]
	}

	// fmt.Println(A, B, num, min, mod)

	for i := 0; i < N; i++ {
		if A[i] <= num {
			fmt.Println(min + 1)
		} else {
			fmt.Println(min)
		}
	}

}
