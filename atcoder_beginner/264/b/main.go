package main

import "fmt"

func main() {

	var R, C int
	fmt.Scan(&R, &C)

	pt := make([]bool, 15)
	for i := 0; i < 15; i++ {
		pt[i] = true
	}
	pf := make([]bool, 15)

	p1 := pt
	p3 := make([]bool, 15)
	copy(p3, p1)
	p3[1] = false
	p3[13] = false
	p5 := make([]bool, 15)
	copy(p5, p3)
	p5[3] = false
	p5[11] = false
	p7 := make([]bool, 15)
	copy(p7, p5)
	p7[5] = false
	p7[9] = false

	p2 := pf
	p2[0] = true
	p2[14] = true
	p4 := make([]bool, 15)
	copy(p4, p2)
	p4[2] = true
	p4[12] = true
	p6 := make([]bool, 15)
	copy(p6, p4)
	p6[4] = true
	p6[10] = true
	p8 := make([]bool, 15)
	copy(p8, p6)
	p8[6] = true
	p8[8] = true

	P := [][]bool{p1, p2, p3, p4, p5, p6, p7, p8, p7, p6, p5, p4, p3, p2, p1}

	res := P[R-1][C-1]
	if res {
		fmt.Println("black")
	} else {
		fmt.Println("white")
	}
}
