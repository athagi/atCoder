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

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	S := make([]int, N)
	T := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		s, _ := strconv.Atoi(sc.Text())
		S[i] = s
	}
	for i := 0; i < N; i++ {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		T[i] = t
	}

	res := make([]int, N)
	copy(res, T)

	counter := 0
	res[0] = T[0]
	for i := 0; ; {
		if i == N-1 {
			if res[i]+S[i] <= res[0] {
				res[0] = res[i] + S[i]
			}
		} else {
			if res[i]+S[i] <= res[i+1] {
				res[i+1] = res[i] + S[i]
			}
		}

		if i == N-1 {
			if counter > 0 {
				break
			}
			i = 0
			counter++
		} else {
			i++
		}

		if counter == N {
			break
		}
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
