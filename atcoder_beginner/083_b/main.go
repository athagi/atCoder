package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, n int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	nums := make([]int, 0)

	for i := 0; i <= n; i++ {
		strNum := strconv.Itoa(i)
		num := 0
		for j := 0; j < len(strNum); j++ {
			strDigit := strNum[j : j+1]
			digit, _ := strconv.Atoi(strDigit)
			num = num + digit
		}

		if num >= a && num <= b {
			nums = append(nums, i)
		}
	}
	res := 0
	for _, i := range nums {
		res = res + i
	}
	fmt.Println(res)
}
