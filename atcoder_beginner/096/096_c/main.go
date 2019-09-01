package main

import (
	"fmt"
	"unicode/utf8"
)

var mark = "#"

func main() {
	var H, W int
	fmt.Scan(&H)
	fmt.Scan(&W)

	inputs := scanStrings(H)
	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	res := true
	for i, s := range inputs {
		for j, ss := range s {
			if ss != mark {
				continue
			}
			check := false
			for k := 0; k < 4; k++ {
				if (i+dy[k] >= 0 && i+dy[k] < H) && (j+dx[k] >= 0 && j+dx[k] < W) {
					if inputs[i+dy[k]][j+dx[k]] == mark {
						check = true
						break
					}
				}
			}
			if !check {
				res = false
			}

		}
		if !res {
			break
		}
	}

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func scanStrings(len int) (strings [][]string) {
	var str string
	for i := 0; i < len; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, spritString(str))
	}
	return
}

func spritString(input string) (res []string) {
	for i := 0; i < utf8.RuneCountInString(input); i++ {
		res = append(res, input[i:i+1])
	}
	return
}
