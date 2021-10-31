package main

import "fmt"

func main() {
	var H, W int
	var A string
	fmt.Scan(&H, &W)

	counter := 0
	board := make([][]string, H)
	for i := 0; i < H; i++ {

		fmt.Scan(&A)
		tmp := make([]string, W)
		for j := 0; j < W; j++ {
			a := A[j : j+1]
			tmp[j] = a
			if a == "#" {
				counter++
			}
		}
		board[i] = tmp
	}

	res := "Possible"
	if counter != W+H-1 {
		res = "Impossible"
	}

	fmt.Println(res)
}
