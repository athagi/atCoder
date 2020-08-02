package main

import "fmt"

func main() {
	var A,B,C,K int
	fmt.Scan(&A, &B,&C,&K)

	var a,b,c int
	if K >= A {
		a = A
		K = K - a
	} else {
		a = K
		K = 0
	}

	if K >= B {
		b = B
		K = K -b
	} else {
		b = K
		K = 0
	}

	if K == C {
		c = C
	} else {
		c = K

	}

	res := a -c


	fmt.Println(res)

}
