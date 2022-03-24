package main

import "fmt"

func main() {
	var A, B, C, N int
	fmt.Scan(&N, &A, &B, &C)

	res := 10000
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000-i; j++ {

			sum := i*A + j*B
			if sum > N {
				break
			}

			if (N-sum)%C == 0 {
				tmp := i + j + ((N - sum) / C)
				if tmp < res {
					res = tmp
				}
			}

		}
	}
	fmt.Println(res)
}
