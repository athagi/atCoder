package main

import (
	"fmt"
)

func main() {
	var H, W, X, Y int
	fmt.Scan(&H, &W, &X, &Y)
	var s string

	S := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&s)
		S[i] = s
	}

	res := 1
	for i := Y - 2; i >= 0; i-- {
		if S[X-1][i:i+1] == "." {
			res++
		} else {
			break
		}
	}

	for i := Y; i < W; i++ {
		if S[X-1][i:i+1] == "." {
			res++
		} else {
			break
		}
	}

	for i := X - 2; i >= 0; i-- {
		if S[i][Y-1:Y] == "." {
			res++
		} else {
			break
		}
	}

	for i := X; i < H; i++ {
		if S[i][Y-1:Y] == "." {
			res++
		} else {
			break
		}
	}

	fmt.Println(res)

}
