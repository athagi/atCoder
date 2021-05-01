package main

import (
		"fmt"
)



func main() {
	var S,s string
	fmt.Scan(&S)

	rev := false
	res := []string{}
	for i:= 0; i < len(S); i++ {
		fmt.Println(S[i:i+1])
		s = S[i:i+1]
		if s == "R" {
			rev = !rev
			continue
		}

		if rev {
			tmp := []string{s}
			res = append(tmp, res...)
		} else {
			res = append(res, s)
		}
	}


	if rev {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	fmt.Println(res)

}

