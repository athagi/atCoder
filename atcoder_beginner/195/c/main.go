package main

import (
		"fmt"
)



func main() {
	var N int64

	fmt.Scan(&N)

	var ans int64
	if N >= 1000 {
		ans += N - 999
	}
	if N >= 1000000{
		ans += N - 999999
	}
	if N >= 1000000000{
		ans += N - 999999999
	}
	if N >= 1000000000000{
		ans += N - 999999999999
	}
	if N >= 1000000000000000{
		ans += N - 999999999999999
	}

	fmt.Println(ans)
}
