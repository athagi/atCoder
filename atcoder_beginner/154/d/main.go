package main

import (
	"fmt"
	"math"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	exceptedValues := make([]float64, N+1)
	var exceptedValue float64
	input := 0
	var res float64
	for i := 0; i < N; i++ {
		fmt.Scan(&input)
		exceptedValue = getExceptedValue(input)

		if i == 0 {
			exceptedValues[i] = exceptedValue
		} else {
			exceptedValues[i] = exceptedValue + exceptedValues[i-1]
		}
		if i+1 >= K {
			if i+1-K != 0 {
				res = math.Max(res, exceptedValues[i]-exceptedValues[i-K])
			} else {
				res = math.Max(res, exceptedValues[i])
			}
		}
	}
	fmt.Println(res)
}

func getExceptedValue(N int) float64 {
	cont := 0
	for i := 0; i < N; i++ {
		cont += i + 1
	}
	res := float64(cont) / float64(N)
	return res
}
