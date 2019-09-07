package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	inputs := scanNums(N)
	ans := 0
	for i := N - 2; i >= 0; i-- {

		ans += inputs[i]
		// fmt.Println(ans)
	}
	ans += inputs[0]
	fmt.Println(ans)
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
