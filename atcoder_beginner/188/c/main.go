package main

import (
    "bufio"
    "fmt"
    "os"
		"strconv"
		"math"
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

type Player struct {
	score int
	position int
}

func main() {
    sc.Split(bufio.ScanWords)
		n := nextInt()
		N := int(math.Pow(2, float64(n)))
		// fmt.Println(N)
		// fmt.Println(n)
		tmp := 0
		p1 := Player{}
		p2 := Player{}

    for i := 0; i < N/2; i++ {
			tmp = nextInt()
			// fmt.Println(tmp)
			if p1.score < tmp {
				p1.score = tmp
				p1.position = i+1
			} 
		}
		// fmt.Println(p1)
		// fmt.Println(p2)

		for i := N/2; i < N; i++ {
			tmp = nextInt()
			if p2.score < tmp {
				p2.score = tmp
				p2.position = i+1
			} 
		}
		
		res := p1.position
		if p1.score > p2.score {
			res = p2.position
		}

		// fmt.Println(p1)
		// fmt.Println(p2)
    fmt.Println(res)
}