package main

import (
	"fmt"
	"sort"
)

type item struct {
	city  string
	point int
	num   int
}

func main() {
	var N, p int
	var s string
	fmt.Scan(&N)

	res := make([]item, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s, &p)
		res[i] = item{s, p, i + 1}
	}
	sort.SliceStable(res, func(i, j int) bool { return res[i].point > res[j].point })
	sort.SliceStable(res, func(i, j int) bool { return res[i].city < res[j].city })

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].num)
	}

}
