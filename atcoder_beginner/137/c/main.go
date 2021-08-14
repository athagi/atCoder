package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	words := make([]string, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		sc.Scan()
		s := sc.Text()
		words[i] = sortString(s)
	}
	sort.Strings(words)

	m := map[string]int{}

	for i := 0; i < len(words); i++ {
		word := words[i]
		if _, ok := m[word]; ok {
			m[word] = m[word] + 1
		} else {
			m[word] = 1
		}
	}

	res := 0
	for k, v := range m {
		fmt.Println(k, v)
		if v > 0 {
			res += calcConv(v)
		}
	}
	fmt.Println(res)

}

func calcConv(n int) int {
	return n * (n - 1) / 2

}

func sortString(word string) string {
	chars := []string{}
	for _, c := range word {
		chars = append(chars, string(c))
	}
	// sort.Sort(sort.StringSlice(chars))
	sort.Strings(chars)
	return strings.Join(chars, "")
}
