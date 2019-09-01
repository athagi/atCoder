package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	inputs := scanNums(N)

	topRecord := 0
	currentRecord := 0
	for i := 0; i < N-1; i++ {
		if inputs[i] >= inputs[i+1] {
			currentRecord++
		} else {
			currentRecord = 0
		}
		if currentRecord > topRecord {
			topRecord = currentRecord
		}
	}
	fmt.Println(topRecord)
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
