package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var a, b, N int
	fmt.Scan(&N)

	A := make([]int, N)
	B := make([]int, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		A[i] = a
	}
	for i := 0; i < N; i++ {
		sc.Scan()
		b, _ = strconv.Atoi(sc.Text())
		B[i] = b
	}

	sort.Ints(A)
	sort.Ints(B)

	res := 0
	for i := 0; i < N; i++ {
		tmp := A[i] - B[i]
		if tmp < 0 {
			tmp = tmp * -1
		}
		res += tmp
	}
	fmt.Println(res)

}
