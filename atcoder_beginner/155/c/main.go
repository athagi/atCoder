package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	var s string
	dict := map[string]int{}
	maxCount := 1
	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		if _, ok := dict[s]; ok {
			dict[s] = dict[s] + 1
			if maxCount < dict[s] {
				maxCount = dict[s]
			}
		} else {
			dict[s] = 1
		}
	}
	slice := []string{}
	for key, val := range dict {
		if val == maxCount {
			slice = append(slice, key)

		}
	}
	sort.Strings(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

}
