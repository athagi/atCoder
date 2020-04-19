package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)
	var res int
	for i := 1; i <= K; i++ {
		for j := 1; j <= K; j++ {
			for k := 1; k <= K; k++ {
				res += gcd(i, gcd(j, k))
			}
		}
	}
	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
