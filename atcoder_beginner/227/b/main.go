package main

import "fmt"

func main() {
	var N, S int
	fmt.Scan(&N)
	possible := make([]int, 1000)
	i := 0
	tmp := 0
	for a := 1; a < 150; a++ {
		for b := a; b < 150; b++ {
			tmp = 4*a*b + 3*a + 3*b
			if tmp <= 1000 {
				possible[i] = tmp
				i++
			} else {
				continue
			}
		}
	}
	// fmt.Println(possible)

	res := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&S)
		flag := false
		for j := 0; j < len(possible); j++ {
			if S == possible[j] {
				flag = true
			}
		}
		if !flag {
			res++
		}
	}
	fmt.Println(res)

}
