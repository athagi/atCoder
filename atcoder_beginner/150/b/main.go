package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)
	counter := 0
	for i := 0; i < N-2; i++ {
		char := S[i : i+1]
		if char == "A" {
			if S[i+1:i+2] == "B" && S[i+2:i+3] == "C" {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
