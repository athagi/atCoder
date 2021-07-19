package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, res int64
	fmt.Scan(&a, &b, &c)

	res = math.MaxInt64
	if a == b && b == c && c == a {
		res = 0
	} else if a == b || b == c || c == a {

		min := math.Min(float64(a), math.Min(float64(b), float64(c)))
		minCount := 0
		if min == float64(a) {
			minCount++
		}
		if min == float64(b) {
			minCount++
		}
		if min == float64(c) {
			minCount++
		}

		x := math.Abs(float64(a) - float64(b))
		y := math.Abs(float64(b) - float64(c))
		z := math.Abs(float64(c) - float64(a))
		add := int64(math.Max(x, math.Max(y, z)))

		if minCount == 2 {
			res = add * 2
		} else {
			res = add
		}

	} else {
		var counter int64
		// center
		if a > b && c > b {
			min := math.Min(float64(a), float64(c))
			counter += int64(min) - b
			b += counter
		}
		// fmt.Println(a, b, c)
		if (a+c)%2 == 1 {
			if a < c {
				a++
				counter++
			} else {
				c++
				counter++
			}
		}
		// fmt.Println(a, b, c)
		x := b - a
		y := c - b
		// add a or c number
		counterA := int64(math.Abs(float64(x) - float64(y)))
		// add b number
		counterB := (a+c)/2 - b

		if counterA < 0 {
			counter += counterB
		} else if counterB < 0 {
			counter += counterA
		} else {
			counter += int64(math.Min(float64(counterA), float64(counterB)))
		}
		res = counter

	}
	fmt.Println(res)
}
