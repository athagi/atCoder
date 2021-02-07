package main

import (
		"fmt"
)

func main() {
	var M, N, A, B, K, C, D int

	fmt.Scan(&N, &M)
	dishes := make([][]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&A, &B)
		dishes[i] = []int{A, B}
	}

	fmt.Scan(&K)
	man := make([][]int, K)
	for i := 0; i < K; i++{
		fmt.Scan(&C, &D)
		man[i] = []int{C, D}
	}

	fill := [][]bool{}
	for bits := 0; bits <(1 << uint64(K)); bits++ {
		f := make([]bool, N+1)
		for i := 0; i< K; i++ {
			if(bits >> uint64(i))&1 == 1 {
				f[man[i][1]] = true
			} else {
				f[man[i][0]] = true
			}
		}
		fill = append(fill, f)
	}

	res := 0
	for _, f := range fill {
		cnt := 0
		for j := 0; j < M; j++ {
			if f[dishes[j][0]] && f[dishes[j][1]] {
				cnt++
			}
		}
		if cnt > res {
			res = cnt
		}
	}

	fmt.Println(res)
}
