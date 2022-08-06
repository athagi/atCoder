package main

import (
	"fmt"
	"sort"
)

func main() {
	var A, B, C, D, E int
	fmt.Scan(&A, &B, &C, &D, &E)
	card := []int{A, B, C, D, E}
	sort.Sort(sort.IntSlice(card))

	res := "No"
	if card[0] == card[1] && card[1] == card[2] && card[2] == card[0] && card[3] == card[4] && card[2] != card[3] {
		res = "Yes"
	}
	if card[2] == card[3] && card[3] == card[4] && card[4] == card[2] && card[0] == card[1] && card[1] != card[2] {
		res = "Yes"
	}

	fmt.Println(res)

}
