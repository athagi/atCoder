package main

import "fmt"

func main() {
	var A,B,C,D int
	fmt.Scan(&A, &B,&C,&D)
	res := "No"
	for {
		C = C - B
		if C <= 0{
			res = "Yes"
			break
			
		}
		A = A - D
		if A <= 0 {
			break
		}
	}

	fmt.Println(res)

}
