package main

import (
	"fmt"
)

func main() {
	var S1, S2, S3 string
	fmt.Scan(&S1, &S2, &S3)

	inputs := []string{S1, S2, S3}
	contests := []string{"ABC", "ARC", "AGC", "AHC"}
	for i := 0; i < len(inputs); i++ {
		contests = remove(contests, inputs[i])
	}

	fmt.Println(contests[0])

}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
