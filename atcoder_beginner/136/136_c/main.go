package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	inputs := scanNums(N)
	res := true
	for i := 0; i < len(inputs)-1; i++ {
		if inputs[i] == inputs[i+1] {
			continue
		} else if inputs[i] <= inputs[i+1]-1 {
			inputs[i+1] = inputs[i+1] - 1
		} else {
			res = false
		}
	}
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
