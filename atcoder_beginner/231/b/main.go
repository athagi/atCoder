package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	s := make(map[string]int)
	res := 0
	resName := ""
	for i := 0; i < N; i++ {
		fmt.Scan(&S)
		_, hasKey := s[S]
		if !hasKey {
			s[S] = 1
		} else {
			s[S] = s[S] + 1
		}
		if s[S] > res {
			res = s[S]
			resName = S
		}
	}

	fmt.Println(resName)
}
