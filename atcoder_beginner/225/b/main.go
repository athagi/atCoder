package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	P := make(map[int]int)
	for i := 0; i < 2*N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if _, ok := P[a]; ok {
			P[a] = P[a] + 1
		} else {
			P[a] = 1
		}
	}
	countOne := 0

	res := "Yes"
	for i := 1; i < len(P); i++ {
		if P[i] == 1 {
			countOne++
		} else {
			if P[i] != N-1 {
				res = "No"
			}
		}
	}
	if countOne != N-1 {
		res = "No"
	}

	fmt.Println(res)
}
