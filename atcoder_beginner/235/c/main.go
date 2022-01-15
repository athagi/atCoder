package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, Q, a, x, k int
	fmt.Scan(&N, &Q)
	m := make(map[int][]int, N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		if _, ok := m[a]; ok {
			m[a] = append(m[a], i+1)
		} else {
			m[a] = []int{i + 1}
		}
	}

	for i := 0; i < Q; i++ {
		sc.Scan()
		x, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		k, _ = strconv.Atoi(sc.Text())
		if len(m[x]) >= k {
			fmt.Println(m[x][k-1])
		} else {
			fmt.Println(-1)
		}
	}
}
