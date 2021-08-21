package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	res := "WA"

	if S == "Hello,World!" {
		res = "AC"
	}

	fmt.Println(res)

}
