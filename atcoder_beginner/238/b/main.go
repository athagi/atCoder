package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, a int
	fmt.Scan(&N)

	pizza := make([]int, N+2)
	pizza[0] = 0
	pizza[1] = 360
	sum := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		sum += a
		if sum >= 360 {
			sum = sum - 360
		}
		pizza[i+2] = sum
	}

	sort.Sort(sort.IntSlice(pizza))
	// fmt.Println(pizza)

	res := 0
	tmp := 0
	for i := 0; i < len(pizza)-1; i++ {
		tmp = pizza[i+1] - pizza[i]
		if res < tmp {
			res = tmp
		}
	}
	fmt.Println(res)
}
