package main

import (
	"fmt"
)

// 1WA になる。２つのパーツに分かれるパターンが網羅できていない。
// 公式解法でとき直す必要あり

func main() {
	var N, M, A, B, C, D int
	fmt.Scan(&N, &M)

	takahashi := make(map[int]int)
	aoki := make(map[int]int)
	for i := 1; i <= 8; i++ {
		takahashi[i] = 0
		aoki[i] = 0
	}

	for i := 0; i < M; i++ {
		fmt.Scan(&A, &B)
		takahashi[A] = takahashi[A] + 1
		takahashi[B] = takahashi[B] + 1
	}

	takahashiRes := convertMap(takahashi)

	for i := 0; i < M; i++ {
		fmt.Scan(&C, &D)
		aoki[C] = aoki[C] + 1
		aoki[D] = aoki[D] + 1
	}
	aokiRes := convertMap(aoki)
	fmt.Println("takahashi=", takahashi)
	fmt.Println("aoki=", aoki)

	fmt.Println("takahashiRes=", takahashiRes)
	fmt.Println("aokiRes=", aokiRes)
	res := "Yes"
	for i := 0; i < len(takahashiRes); i++ {
		fmt.Println(takahashiRes[i], aokiRes[i], i)
		if takahashiRes[i] != aokiRes[i] {
			res = "No"
			// break
		}
	}
	fmt.Println(res)
}

func convertMap(in map[int]int) map[int]int {
	res := make(map[int]int)
	for i := 0; i < len(in); i++ {
		res[i] = 0
	}
	for i := 1; i <= len(in); i++ {
		tmp := in[i]
		res[tmp] = res[tmp] + 1
	}
	return res
}
