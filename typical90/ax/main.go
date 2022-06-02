package main

import "fmt"

func main() {
	var N, L int
	fmt.Scan(&N, &L)

	res := make([]int, N+1)
	res[0] = 1

	for i := 0; i < N+1; i++ {
		if i+1 <= len(res)-1 {
			tmp := (res[i+1] + res[i]) % 1000000007
			res[i+1] = tmp
		}
		if i+L <= len(res)-1 {
			tmp := (res[i+L] + res[i]) % 1000000007
			res[i+L] = tmp
		}

	}
	fmt.Println(res[len(res)-1])
}
