package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

func main() {
	var N,M,K int

	
	fmt.Scan(&N,&M,&K)

	As := make([]int, N+1)
	Bs := make([]int, M+1)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var sum int
	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sum += a
		As[i+1] = sum		
	}

	sum = 0
	counter := 0
	for i := 0; i < M; i++ {
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		sum += b
		if sum > K {
			break
		}
		Bs[i+1] = sum
		counter = i+1
	}
	Bs = Bs[0:counter+1]

	var a, b int
	var max float64
	bLoc := len(Bs) -1
	for i := 0; i < len(As); i++ {
		a = As[i]
		if a > K {
			break
		}
		for j := bLoc; j >= 0; j-- {
			b = Bs[j]
			
			if a + b > K{
				continue
			}
			// fmt.Println(i,j,a,b, a+b)
			max = math.Max(max, (float64)(i + j))
			bLoc = j
			break
		}
	}
	fmt.Println(max)

}

