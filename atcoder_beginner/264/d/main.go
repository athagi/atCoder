package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	dict := map[string]int{"a": 1, "t": 2, "c": 3, "o": 4, "d": 5, "e": 6, "r": 7}
	s := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		tmp := S[i : i+1]
		s[i] = dict[tmp]
	}
	res := BubbleSort(s)
	fmt.Println(res)
}

func BubbleSort(a []int) int {
	cnt := 0
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				cnt++
			}
		}
	}

	return cnt
}
