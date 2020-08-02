package main

import (
	"fmt"
)

func main() {
	var N int 
	fmt.Scan(&N)
	var s string
	var ac, wa, tle, re int
	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		if s == "AC" {
			ac++
		} else if s == "WA" {
			wa++
		} else if s == "TLE" {
			tle++
		} else {
			re++
		}
	}
	fmt.Println("AC x", ac)
	fmt.Println("WA x", wa)
	fmt.Println("TLE x", tle)
	fmt.Println("RE x", re)

	
}

