// これだと 2WA になる
package main

import (
	"fmt"
	"math"
)

func main() {
	var N float64
	fmt.Scan(&N)
	var res int64
	var i int64
	for i = 0; ; i++ {
		if math.Exp2(float64(i)) > N {
			break
		}
		res = i
	}
	fmt.Println(res)
}
