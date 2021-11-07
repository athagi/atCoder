package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var a string
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)

	res := make(map[string]bool)

	for i := 0; i < N; i++ {
		sc.Scan()
		a = sc.Text()

		if !res[a] {
			res[a] = true
		}
	}
	fmt.Println(len(res))
}
