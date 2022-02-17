package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	graph := make([][]int, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < M; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		graph[a-1] = append(graph[a-1], b)
		graph[b-1] = append(graph[b-1], a)
	}
	res := 0
	for i := 0; i < len(graph); i++ {
		a := graph[i]
		cnt := 0
		for j := 0; j < len(a); j++ {
			if a[j] < i+1 {
				cnt++
			}
		}
		if cnt == 1 {
			res++
		}
	}
	fmt.Println(res)
}
