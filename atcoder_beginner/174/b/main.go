package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
	}
	return i
}

func main() {
	var N int
	var D int64
	var x, y int64
	res := 0

    sc.Split(bufio.ScanWords)
	N = nextInt()
	D = int64(nextInt())
	var DD int64
	DD = D * D

    for i := 0; i < N; i++ {
		x = int64(nextInt())
		y = int64(nextInt())
		
		if (DD >= x*x + y*y){
			res++
		}
    }

	fmt.Println(res)
}


   