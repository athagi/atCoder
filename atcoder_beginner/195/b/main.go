package main

import (
	"fmt"

)

func main() {
	var A, B, W int
	fmt.Scan(&A, &B, &W)

	w := W * 1000

	min := 0
	max := 0
	for i := 1;;i++ {
		if A*i <= w && w <= B*i {
			if min == 0 {
				min = i
			}
			continue
		}

		if min > 0 {
			max = i -1
			break
		}

		if A*i > w {
			break
		}
	}

	if min == 0 && max == 0 {
		fmt.Println("UNSATISFIABLE")
	} else {
		fmt.Println(min, max)
	}
}