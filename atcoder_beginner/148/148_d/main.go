package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	res := solv(N)

	fmt.Println(res)

}

func solv(len int) (broken int) {
	counter := 0
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		if num == counter+1 {
			counter += 1
		} else {
			broken += 1
		}
		// nums = append(nums, num)
	}
	if broken == len {
		return -1
	}
	return
}

// func scanNums(len int) (nums []int) {
// 	var num int
// 	var broken int
// 	for i := 0; i < len; i++ {
// 		fmt.Scan(&num)
// 		nums = append(nums, num)
// 	}
// 	return
// }
