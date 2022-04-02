package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var N, K, X, a int
	fmt.Scan(&N, &K, &X)
	A := make([]int, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())

		if a >= X && K > 0 {
			tmp := a / X
			if K >= tmp {
				a = a - tmp*X
				K = K - tmp
			} else {
				a = a - K*X
				K = 0
			}
		}
		A[i] = a
	}
	sort.Sort(sort.IntSlice(A))
	if K > 0 {
		tmp := 0
		if N-K > 0 {
			tmp = N - K
		}
		A = A[:tmp]
	}

	res := 0
	for i := 0; i < len(A); i++ {
		res += A[i]
	}
	fmt.Println(res)
}
