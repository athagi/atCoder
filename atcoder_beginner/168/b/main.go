package main

import "fmt"

func main() {
	var K int
	var S string
	fmt.Scan(&K, &S)

	res := S
	if len(S) > K {
		res = S[0:K] + "..."
	} 

	fmt.Println(res)

}
