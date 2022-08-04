package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	cake := make([]int, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	size := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		cake[i] = a
		size += a
	}

	res := "No"
	if size%10 == 0 {
		target := size / 10
		tmp := 0
		start := 0
		end := 0
		flag := false
		for {
			if tmp < target {
				start++
				if start > end && flag {
					break
				}
				if start > len(cake)-1 {
					start = 0
					flag = true
				}
				tmp += cake[start]
			} else if tmp > target {
				end++
				if end > len(cake)-1 {
					break
				}
				tmp -= cake[end]
			}
			// fmt.Println(tmp, target)
			if tmp == target {
				res = "Yes"
				break
			}

		}

	}
	fmt.Println(res)
}
