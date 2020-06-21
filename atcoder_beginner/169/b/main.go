package main

import (
	"fmt"
	"os"
)

func main() {
	var N int 
	var A int64
	fmt.Scan(&N)
	As := make([]int64, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		As[i] = A
		if A == 0 {
			fmt.Println(0)
			os.Exit(0)
		}
	}


	var res int64
	res = 1
	for i := 0; i < N; i++ {
		n := As[i]
		res = res * n
		if res == -1 {
			break;
		}
		if res > 1000000000000000000 || res < 0 {
			res = -1
			break;
		}

	}
	
	fmt.Println(res)

}

// func add(a,b int64) (res int64) {
	
//     defer func() {
//         err := recover()
//         if err != nil {
// 			fmt.Println("aa")
//             res = -1
// 		}
//     }()
	
// 	res = a*b
// 	return res
// }
