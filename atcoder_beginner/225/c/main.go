package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	res := "Yes"
	B := make([][]int, N)
	for i := 0; i < N; i++ {
		tmp := make([]int, M)
		for j := 0; j < M; j++ {
			sc.Scan()
			b, _ := strconv.Atoi(sc.Text())
			tmp[j] = b

			if j > 0 {
				if tmp[j-1]+1 != b {
					res = "No"
				}

				if (tmp[j-1]-1)/7 != (b-1)/7 {
					res = "No"
				}
			}
		}
		B[i] = tmp
	}
	for i := 1; i < len(B); i++ {
		if B[i-1][0]+7 != B[i][0] {
			res = "No"
		}
	}

	fmt.Println(res)
}
