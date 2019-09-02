package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	inputs := scanNums(N)

	maxValue := 0
	secondValue := 0
	maxValueCols := []int{}
	for i, v := range inputs {
		if maxValue == v {
			maxValueCols = append(maxValueCols, i)
		} else if v > maxValue {
			secondValue = maxValue
			maxValue = v
			maxValueCols = []int{i}
		} else if v > secondValue {
			secondValue = v
		} else {
			continue
		}
	}
	res := createRes(N, maxValue, secondValue, maxValueCols)

	for _, v := range res {
		fmt.Println(v)
	}
}

func createRes(N, maxValue, secondValue int, maxValueCols []int) (res []int) {
	for i := 0; i < N; i++ {
		res = append(res, maxValue)
	}
	if len(maxValueCols) == 1 {
		res[maxValueCols[0]] = secondValue
	}
	return
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
