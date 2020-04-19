package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	var A int
	homework := 0
	for i := 0; i < M; i++ {
		fmt.Scan(&A)
		homework += A
	}
	res := N - homework
	if res < 0 {
		res = -1
	}

	fmt.Println(res)

}
