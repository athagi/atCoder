package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	n := len(S)

	t := " chokudai"
	dp := make([][]int, len(t))

	banpei := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		banpei[i] = 1
	}
	dp[0] = banpei
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	mod := 1000000007

	for i := 1; i < len(t); i++ {
		for j := 0; j < n; j++ {
			if t[i:i+1] != S[j:j+1] {
				dp[i][j+1] = dp[i][j]
			} else {
				dp[i][j+1] = (dp[i][j] + dp[i-1][j]) % mod
			}
		}
	}
	fmt.Println(dp[len(dp)-1][n])
}
