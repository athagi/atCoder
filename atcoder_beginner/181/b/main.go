package main

import (
	"fmt"
)



func main() {
	var N int
	var A, B,res,tmp int64
	fmt.Scan(&N)
	
	
	for i := 0; i < N; i++ {
		fmt.Scan(&A,&B)
		tmp = B-A+1
		if tmp %2 == 0 {
			res += (B+A)*tmp/2
		} else {
			res += ((B+A)*(tmp-1))/2 + (B+A)/2
		}
	}

	fmt.Println(res)

}
