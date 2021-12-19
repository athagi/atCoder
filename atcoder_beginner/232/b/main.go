package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)

	res := "Yes"
	distance := calc(S[0:1], T[0:1])
	for i := 0; i < len(S); i++ {
		a := S[i : i+1]
		b := T[i : i+1]
		tmp := calc(a, b)
		// fmt.Println(distance, tmp)
		if distance != tmp {
			res = "No"
			break
		}
	}
	fmt.Println(res)

}

func calc(a, b string) int {
	ALPHABET := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var A, B int
	for i := 0; i < len(ALPHABET); i++ {
		if a == ALPHABET[i] {
			A = i
		}
		if b == ALPHABET[i] {
			B = i
		}
	}

	res := B - A
	if res < 0 {
		res = res + len(ALPHABET)
	}
	return res
}
