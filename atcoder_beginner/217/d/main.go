// WIP
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var L, Q int
	fmt.Scan(&L, &Q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	C := make([]int, Q)
	X := make([]int, Q)
	for i := 0; i < Q; i++ {
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		C[i] = c
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		X[i] = x
	}

	wood := []int{0, L, 0}
	for i := 0; i < Q; i++ {
		pos := 0
		tmp1 := 0
		for j := 0; j < len(wood); j++ {
			if tmp1 < X[i] && X[i] < tmp1+wood[j] {
				pos = j
				break
			}
			tmp1 += wood[j]
		}

		if C[i] == 1 {
			a := wood[:pos]
			b := make([]int, len(wood[pos+1:]))
			copy(b, wood[pos+1:])

			c := append(a, []int{X[i] - tmp1, wood[pos] - (X[i] - tmp1)}...)
			c = append(c, b...)
			wood = c
		} else {
			fmt.Println(wood[pos])
		}
		// fmt.Println("--", wood)
	}

}
