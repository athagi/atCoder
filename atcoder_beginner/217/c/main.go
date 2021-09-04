package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	P := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		P[i] = p
	}

	Q := make([]string, N)
	for i := 0; i < N; i++ {
		Q[P[i]-1] = strconv.Itoa(i + 1)
	}
	fmt.Println(strings.Join(Q, " "))
}
