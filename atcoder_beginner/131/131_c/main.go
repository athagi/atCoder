package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	cStart := (a - 1) / c
	cEnd := b / c
	cNum := cEnd - cStart

	dStart := (a - 1) / d
	dEnd := b / d
	dNum := dEnd - dStart

	cdNum := 0
	if cNum > 0 && dNum > 0 {
		minCommonMultiple := c * d / getMinCommonDivisor(c, d)

		cdStart := (a - 1) / minCommonMultiple
		cdEnd := b / minCommonMultiple

		cdNum = cdEnd - cdStart

	}
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
