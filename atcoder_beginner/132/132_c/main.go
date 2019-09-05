package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

func main() {
	var N int
	fmt.Scan(&N)

	inputs := scanNums(N)

	sort.Sort(sort.IntSlice(inputs))
	a := inputs[N/2-1]
	b := inputs[N/2]

	fmt.Println(b - a)

}

func scanStrings(len int) (strings []string) {
	var str string
	for i := 0; i < len; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}
	return
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func spritString(input string) (res []string) {
	for i := 0; i < utf8.RuneCountInString(input); i++ {
		res = append(res, input[i:i+1])
	}
	return
}
