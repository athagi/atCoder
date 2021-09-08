package main

import "fmt"

func main() {
	var S string
	var N int
	fmt.Scan(&N, &S)

	res := 0
	for i := 0; i < N; i++ {
		A := S[:i]
		B := S[i:]
		X := make(map[string]bool)
		Y := make(map[string]bool)
		for j := 0; j < len(A); j++ {
			if !X[A[j:j+1]] {
				X[A[j:j+1]] = true
			}
		}
		for j := 0; j < len(B); j++ {
			if !Y[B[j:j+1]] {
				Y[B[j:j+1]] = true
			}
		}

		cnt := 0
		for k, _ := range X {
			if Y[k] {
				cnt++
			}
		}

		if cnt > res {
			res = cnt
		}

	}
	fmt.Println(res)
}
