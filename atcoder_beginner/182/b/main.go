package main

import (
	"fmt"
)

func main() {
	var N, A int
	fmt.Scan(&N)
	As := make([]int, N)
	
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		As[i] = A
	}
	
	res := 0
	resNum := 0
	count :=0
	for i := 2; i <= 1000; i++ {
		for j := 0; j < len(As); j++ {
			if As[j] % i == 0 {
				count++
			}
		}
		
		if res < count {
			res = count
			resNum = i
		}
		count = 0
	}
	fmt.Println(resNum)
}