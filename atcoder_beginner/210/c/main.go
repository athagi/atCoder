package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Candy struct {
	Number int
	count  int
}

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	C := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		C[i] = a
	}

	k := C[0:K]
	m := make(map[int]int)
	for _, ele := range k {
		if _, ok := m[ele]; !ok {
			m[ele] = 1
		} else {
			m[ele] += 1
		}
	}

	count := len(m)
	for i := 0; i < N-K; i++ {
		removeNum := C[i]
		// fmt.Println(removeNum)
		if m[removeNum] == 1 {
			delete(m, removeNum)
		} else {
			m[removeNum] -= 1
		}

		appendNum := C[i+K]
		// fmt.Println(appendNum)
		if m[appendNum] > 0 {
			m[appendNum] += 1
		} else {
			m[appendNum] = 1
		}

		if count <= len(m) {
			count = len(m)
		}
	}
	fmt.Println(count)
}
