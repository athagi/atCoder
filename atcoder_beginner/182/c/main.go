package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	var A int64
	fmt.Scan(&A)
	
	num := make([]int, len(strconv.FormatInt(A, 10)))
	tmp := A
	for i := 0; i < len(strconv.FormatInt(A, 10)); i++ {
		num[i] = int(tmp % 10)
		tmp = tmp / 10
	}

	sum := 0
	existOne := false
	existTwo := false
	for i := 0; i < len(num); i++ {
		sum += num[i]
		if num[i] % 3 == 1 {
			existOne = true
		}
		if num[i] % 3 == 2 {
			existTwo = true
		}
	}

	res := math.MaxInt16

	switch sum % 3 {
	case 0:
		res = 0
	case 1:
		if existOne {
			res = 1
			if len(num) <= 1 {
				res = -1
			}
		} else {
			res = 2
			if len(num) <= 2 {
				res = -1
			}
		}

	case 2:
		if existTwo {
			res = 1
			if len(num) <= 1 {
				res = -1
			}
		} else {
			res = 2
			if len(num) <= 2 {
				res = -1
			}
		}
	}

	fmt.Println(res)
}

