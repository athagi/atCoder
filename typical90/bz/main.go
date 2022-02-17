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
	graph := make([]int, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < M; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		if graph[a-1] == 0 && a > b {
			graph[a-1] = b
		} else if graph[a-1] >= 1 && a > b {
			graph[a-1] = -1
		}
		if graph[b-1] == 0 && a < b {
			graph[b-1] = a
		} else if graph[b-1] >= 1 && a < b {
			graph[b-1] = -1
		}

	}
	res := 0
	for i := 0; i < len(graph); i++ {
		if graph[i] >= 1 && graph[i] <= i+1 {
			res++
		}
	}
	fmt.Println(res)
}
