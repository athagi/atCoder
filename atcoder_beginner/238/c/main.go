package main

import (
	"fmt"
)

var mod = 998244353

// mod = 998244353 における 2 の逆元は inv2(Inverse element) = 499122177
//https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#3-2-mod-p-%E3%81%AB%E3%81%8A%E3%81%91%E3%82%8B%E9%80%86%E5%85%83
var inv2 = 499122177

func main() {
	var N int64
	fmt.Scan(&N)

	var res, l, r int64
	var p10 int64
	p10 = 10

	for dg := 1; dg <= 18; dg++ {
		l = p10 / 10
		r = N

		// 9, 99, 999 毎に
		if r > p10-1 {
			r = p10 - 1
		}

		// 1桁目, 2桁目... と合計値を足していく（都度mod で割る）
		if l <= r {
			res += calc(r - l + 1)
			res %= int64(mod)

		}
		p10 *= 10
	}
	fmt.Println(res)

}

func calc(n int64) int64 {
	x := n % int64(mod)

	res := x
	res *= (x + 1)
	res %= int64(mod)
	res *= int64(inv2)
	res %= int64(mod)
	return res
}
