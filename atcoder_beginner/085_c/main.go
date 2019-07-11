package main

import (
	"fmt"
	"os"
)

const (
	a = 10000
	b = 5000
	c = 1000
)

func main() {
	var n, y int
	fmt.Scan(&n)
	fmt.Scan(&y)

	for i := 0; i <= n; i++ {
		if a*i > y {
			break
		}
		for j := 0; j <= n-i; j++ {
			if a*i+b*j > y {
				break
			}
			k := n - i - j
			if i*a+j*b+k*c == y {
				fmt.Printf("%d %d %d\n", i, j, k)
				os.Exit(0)
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
