package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type cheese struct {
	A int
	B int
}

func main() {
	var N, W, A, B int
	fmt.Scan(&N, &W)

	cheeses := make([]cheese, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < N; i++ {
		sc.Scan()
		A, _ = strconv.Atoi(sc.Text())

		sc.Scan()
		B, _ = strconv.Atoi(sc.Text())
		cheeses[i] = cheese{A, B}
	}
	sort.Slice(cheeses, func(i, j int) bool {
		return cheeses[i].A > cheeses[j].A
	})
	res := 0
	weight := 0
	for i := 0; i < len(cheeses); i++ {
		cheese := cheeses[i]
		flag := false
		for j := 0; j < cheese.B; j++ {
			if weight+1 > W {
				flag = true
				break
			}
			weight++
			res += cheese.A
		}
		if flag {
			break
		}
	}

	fmt.Println(res)
}
