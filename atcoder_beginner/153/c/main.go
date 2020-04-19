package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	enemies := make([]int, 0)
	enemy := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&enemy)
		enemies = append(enemies, enemy)
	}

	sort.Sort(sort.IntSlice(enemies))
	var res int64
	if len(enemies) <= K {
		res = 0
	} else {
		enemies = enemies[0 : len(enemies)-K]
		for i := 0; i < len(enemies); i++ {
			res += int64(enemies[i])
		}
	}

	fmt.Println(res)
}
