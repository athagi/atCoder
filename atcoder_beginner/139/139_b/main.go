package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	currentTap := 1
	tapCount := 0
	for {
		if b == 1 {
			break
		}
		currentTap = currentTap - 1 + a
		tapCount++
		if currentTap >= b {
			break
		}
	}

	fmt.Println(tapCount)

}
