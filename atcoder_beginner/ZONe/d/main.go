package main

import (
		"fmt"
		"strings"
)



func main() {
	var S,s string
	fmt.Scan(&S)

	rev := false
	res := []string{}
	for i:= 0; i < len(S); i++ {
		// fmt.Println(S[i:i+1])
		s = S[i:i+1]
		if s == "R" {
			rev = !rev
			continue
		}

		if len(res) == 0 {
			res = append(res, s)
			continue
		}

		if rev {
			tmp := []string{s}
			if s == res[0] {
				res = res[1:]
				continue
			}
			res = append(tmp, res...)
		} else {
			if s == res[len(res)-1] {
				res = res[:len(res)-1]
				continue
			}
			res = append(res, s)
		}
	}



	if rev {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	fmt.Println(strings.Join(res, ""))

}