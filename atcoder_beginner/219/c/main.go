package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type input struct {
	word string
	num  string
}

func main() {
	var X string
	var N int
	fmt.Scan(&X, &N)
	X = " " + X

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	S := make([]input, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		s := fmt.Sprintf("%-10s", sc.Text())

		num := ""
		for j := 0; j < len(s); j++ {
			char := s[j : j+1]
			for k := 0; k < len(X); k++ {
				if char == X[k:k+1] {
					a := strconv.Itoa(k)
					b := fmt.Sprintf("%02s", a)
					num += b
				}
			}
		}

		S[i] = input{word: s, num: num}
	}

	sort.SliceStable(S, func(i, j int) bool { return S[i].num < S[j].num })

	for i := 0; i < len(S); i++ {
		fmt.Println(strings.TrimSpace(S[i].word))
	}
}
