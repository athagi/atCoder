package main

import "fmt"

func main() {
	var N, A int
	fmt.Scan(&N)
	res := 1
	for i := 0; i < N; i++ {
		tmp := 0
		for j := 0; j < 6; j++ {
			fmt.Scan(&A)
			tmp += A
		}
		res *= tmp
		res = res % 1000000007
	}
	fmt.Println(res)
}
