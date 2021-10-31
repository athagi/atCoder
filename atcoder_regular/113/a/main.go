package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)

	res := 0
	for i := 1; i <= K; i++ {
		for j := 1; j <= K; j++ {
			if i*j > K {
				break
			}
			res += K / (i * j)
		}
	}
	fmt.Println(res)
}
