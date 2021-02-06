package main

import (
		"fmt"
)



func main() {
	var H, W int
	var s string

	fmt.Scan(&H, &W)
	square := make([][]string, H)
	row := make([]string, W)
	for i := 0; i < H; i++ {
		row = make([]string, W)
		fmt.Scan(&s)
		for j := 0; j < W; j++ {
			row[j] = s[j:j+1]
		}
		square[i] = row
	}

	cnt := 0
	res := 0
	for i := 0; i < H-1; i++ {
		for j := 0; j < W-1; j++ {
			cnt = 0
			if square[i][j] == "#" {
				cnt++
			}
			if square[i+1][j] == "#" {
				cnt++
			}
			if square[i][j+1] == "#" {
				cnt++
			}
			if square[i+1][j+1] == "#" {
				cnt++
			}
			if cnt ==1 || cnt == 3 {
				res++
			}
		}
	}
	

	fmt.Println(res)
}