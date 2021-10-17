package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type input struct {
	word string
	num  string
}

func main() {
	var N int
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var sum float64 // 燃え尽きる秒数
	A := make([]float64, N)
	B := make([]float64, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())

		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		sum += float64(a) / float64(b)
		A[i] = float64(a)
		B[i] = float64(b)
	}

	var res float64
	var target = sum / 2
	for i := 0; i < N; i++ {
		if A[i]/B[i] < target {
			target -= A[i] / B[i]
			res += A[i]
		} else {
			res += B[i] * target
			break
		}

	}
	fmt.Println(res)
}
