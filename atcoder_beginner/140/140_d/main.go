package main

import (
	"fmt"
)

func main() {
	var a, b int
	var c string
	fmt.Scan(&a, &b, &c)

	seat := spritString(c, a)
	ans := 0
	if a == 1 {
		ans = 0
	} else if a == 2 {
		if b > 0 {
			ans = 1
		}
	} else {
		changeCount := b
		for i := 0; i < a-2; i++ {
			if seat[i] == seat[i+2] && seat[i] != seat[i+1] {
				seat[i+1] = seat[i]
				changeCount--
			}
			if changeCount == 0 {
				break
			}
		}

		if changeCount > 0 {
			for i := 1; i < a-1; i++ {
				if seat[i] != seat[i+1] {
					seat[i+1] = seat[i]
				}
			}
		}

		s := seat[0]

		for i := 1; i < a; i++ {
			if s == seat[i] {
				ans++
			}
			s = seat[i]
		}

	}
	fmt.Println(ans)
}

func spritString(input string, n int) []string {
	s := make([]string, n)
	// for i := 0; i < utf8.RuneCountInString(input); i++ {
	// 	res = append(res, input[i:i+1])
	// }
	for i := 0; i < n; i++ {
		s[i] = input[i : i+1]
	}
	return s
}
