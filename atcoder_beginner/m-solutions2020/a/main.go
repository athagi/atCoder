package main

import "fmt"

func main() {
	var X int
	fmt.Scan(&X)

	res := 0
	switch {
	case 400 <= X && X <= 599:
		res = 8
	case 600 <= X && X <= 799:
		res = 7
	case 800 <= X && X <= 999:
		res = 6
	case 1000 <= X && X <= 1199:
		res = 5
	case 1200 <= X && X <= 1399:
		res = 4
	case 1400 <= X && X <= 1599:
		res = 3
	case 1600 <= X && X <= 1799:
		res = 2
	case 1800 <= X && X <= 1999:
		res = 1

	}
	fmt.Println(res)
}
