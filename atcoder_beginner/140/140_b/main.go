package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	inputs := scanNums(2*N + (N - 1))
	A := inputs[0:N]
	B := inputs[N : 2*N]
	C := inputs[2*N:]

	ans := 0
	beforeDish := A[0]
	for i := 0; i < N; i++ {
		ans += B[A[i]-1]
		if beforeDish+1 == A[i] {
			ans += C[A[i-1]-1]
		}

		beforeDish = A[i]
	}
	fmt.Println(ans)

}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
