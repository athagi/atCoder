package main

import (
	"fmt"
	"os"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	inputs := scanNums(M)

	steps := make([]bool, N+1, N+1)

	for i := 0; i <= N; i++ {
		steps[i] = true
	}
	for _, v := range inputs {
		steps[v] = false
		if !steps[v-1] {
			fmt.Println(0)
			os.Exit(0)
		}
	}

	fib := make([]int, N+1, N+1)
	fib[0] = 1
	if steps[1] {
		fib[1] = 1
	} else {
		fib[1] = 0
	}
	for i := 2; i <= N; i++ {
		if steps[i] {
			fib[i] = fib[i-1] + fib[i-2]
		}
		fib[i] %= 1000000007
		// fmt.Println(i, fib[i])
	}

	res := fib[N]
	fmt.Println(res)

}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
