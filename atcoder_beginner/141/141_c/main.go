package main

import (
	"fmt"
)

func main() {
	var N, K, Q int
	fmt.Scan(&N, &K, &Q)
	As := scanNums(Q)
	Points := make([]int, N, N)
	for i := 0; i < N; i++ {
		Points[i] = K
	}

	result := make([]int, N, N)
	for _, v := range As {
		result[v-1]++
	}
	// fmt.Println(result)
	for _, v := range result {
		// fmt.Println(K - Q + v)
		if K-Q+v > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
