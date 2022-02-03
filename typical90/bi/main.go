package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var Q int
	fmt.Scan(&Q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var t, x int
	T := make([]int, Q*2)
	head := Q
	tail := Q
	for i := 0; i < Q; i++ {
		sc.Scan()
		t, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		x, _ = strconv.Atoi(sc.Text())
		if t == 1 {
			T[head-1] = x
			head -= 1
		} else if t == 2 {
			T[tail] = x
			tail += 1
		} else {
			fmt.Println(T[head-1+x])
		}
	}
}
