package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	P := make([]int, N)
	Q := make([]int, N)
	var n int
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		P[i] = n
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		Q[i] = n
	}
	p := resolv(P)
	q := resolv(Q)
	res := p - q
	if p < q {
		res = q - p
	}
	fmt.Println(res)

}

func resolv(n []int) int {
	list := []int{}
	for i := 1; i <= len(n); i++ {
		list = append(list, i)
	}
	res := 1
	for i := 0; i < len(n); i++ {
		char := n[i]
		for j := 0; j < len(list); j++ {
			if char == list[j] {
				res += sum(len(list)) * j / len(list)
				list = remove(list, char)
				break
			}
		}

	}

	return res
}

func sum(n int) int {
	res := 1
	for i := n; i > 1; i-- {
		res *= i
	}
	return res
}

func remove(ints []int, search int) []int {
	result := []int{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}
