package main

import (
	"fmt"
	"sort"
)

func main() {
	var N,K,p int 
	
	fmt.Scan(&N,&K)
	As := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&p)
		As[i] = p
	}

	sort.Sort(sort.IntSlice(As))

	res := 0
	for i := 0; i < K; i++ {
		res += As[i]
	}
	fmt.Println(res)
}

