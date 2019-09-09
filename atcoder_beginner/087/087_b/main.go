package main

import (
	"fmt"
	"os"
)

const (
	x = 500
	y = 100
	z = 50
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	if x*a+y*b+z*c < d {
		fmt.Println("0")
		os.Exit(0)
	}

	counter := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if x*i+y*j+z*k == d {
					counter++
				}
			}
		}
	}
	fmt.Println(counter)

}
