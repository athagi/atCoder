package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	cStart := 0
	cEnd := 0
	cNum := 0
	if c > b {
		cNum = 0
	} else {
		cStart = a / c
		cEnd = b / c
		cNum = cEnd - cStart
	}

	dStart := 0
	dEnd := 0
	dNum := 0

	if d > b {

	} else {
		dStart = a / d
		dEnd = b / d
		dNum = dEnd - dStart
	}

	cdNum := 0
	if cNum > 0 && dNum > 0 {

		minCommonMultiple := c * d / getMinCommonDivisor(c, d)
		if minCommonMultiple <= b {

			cdStart := a / minCommonMultiple
			cdEnd := b / minCommonMultiple

			cdNum = cdEnd - cdStart
		}

	}
	// fmt.Println(c * d / getMinCommonDivisor(c, d))
	// fmt.Println(cNum, dNum, cdNum)
	res := b - a + 1 - (cNum + dNum - cdNum)
	fmt.Println(res)
}

func getMinCommonDivisor(m, n int) int {
	for m%n != 0 {
		temp := n
		n = m % n
		m = temp
	}
	return n
}
