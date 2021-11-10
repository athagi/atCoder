package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, A int
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	nums := make(map[int]int)
	for i := 0; i < 200; i++ {
		nums[i] = 0
	}
	for i := 0; i < N; i++ {
		sc.Scan()
		A, _ = strconv.Atoi(sc.Text())
		if nums[A%200] == 0 {
			nums[A%200] = 1
		} else {
			nums[A%200] = nums[A%200] + 1
		}
	}
	res := 0

	for i := 0; i < 200; i++ {
		if nums[i] > 1 {
			res += combination(nums[i], 2)

		}
	}
	fmt.Println(res)
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}
