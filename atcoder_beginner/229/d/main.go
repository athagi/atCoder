package main

import (
	"fmt"
	"math"
)

func main() {
	var S string
	var K int
	fmt.Scan(&S, &K)

	var res float64
	counterK := K
	length := 0
	for i := 0; i < len(S); i++ {
		if S[i:i+1] == "." {
			if counterK-1 < 0 {
				res = float64(length)
				break
			}
			counterK--
		}
		length++
	}
	if counterK > 0 {
		res = float64(len(S))
	}

	startPosition := 0
	endPosition := length
	fmt.Println(startPosition, endPosition, res, counterK)

	for ; endPosition < len(S); endPosition++ {
		if S[endPosition:endPosition+1] == "." {
			for {
				if S[startPosition:startPosition+1] == "." {
					startPosition++
					res = math.Max(res, float64(endPosition-startPosition+1))
					fmt.Println("+", startPosition, endPosition, res)
					break
				} else {
					startPosition++
				}
			}
		} else {
			res = math.Max(res, float64(endPosition-startPosition+1))
			fmt.Println("-", startPosition, endPosition, res)
		}
	}
	fmt.Println(res)

}
