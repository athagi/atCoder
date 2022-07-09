package main

import (
	"fmt"
)

type Char struct {
	S string
	N int
}

func main() {
	var S, T string
	fmt.Scan(&S, &T)
	s := separateString(S)
	t := separateString(T)

	// fmt.Println(s)
	// fmt.Println(t)
	res := "Yes"
	if len(S) <= len(T) && len(s) == len(t) {
		for i := 0; i < len(t); i++ {
			if s[i].S != t[i].S || s[i].N > t[i].N || (s[i].N == 1 && t[i].N > 1) {
				res = "No"
				break
			}
		}
	} else {
		res = "No"
	}

	fmt.Println(res)
}

func delete_empty(s []Char) []Char {
	var r []Char
	for _, str := range s {
		if str.S != "" {
			r = append(r, str)
		}
	}
	return r
}

func separateString(s string) []Char {
	res := make([]Char, 200000)
	pos := 0
	res[0] = Char{s[0:1], 1}
	for i := 1; i < len(s); i++ {
		tmp1 := res[pos].S[0:1]
		currentChar := s[i : i+1]
		if tmp1 == currentChar {
			if len(tmp1) == 1 {
				res[pos].N = res[pos].N + 1
			}
		} else {
			pos++
			res[pos] = Char{S: currentChar, N: 1}
		}
	}
	return delete_empty(res)
}
