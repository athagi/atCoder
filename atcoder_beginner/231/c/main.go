package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type res struct {
	order int
	input int
	res   int
}

func main() {
	var N, Q, x int
	fmt.Scan(&N, &Q)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var A int
	students := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		A, _ = strconv.Atoi(sc.Text())
		students[i] = A
	}
	sort.Sort(sort.IntSlice(students))
	// fmt.Println(students)

	for i := 0; i < Q; i++ {
		sc.Scan()
		x, _ = strconv.Atoi(sc.Text())
		ok := N
		ng := -1
		var md int
		for ok-ng > 1 {
			md = (ok + ng) / 2
			if students[md] >= x {
				ok = md
			} else {
				ng = md
			}
		}
		fmt.Println(N - ok)

	}

}
