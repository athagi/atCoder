package main

import (
	"fmt"
	"strconv"
	"strings"
)

var bomb = "#"

func main() {
	var H, W int
	fmt.Scan(&H)
	fmt.Scan(&W)

	inputs := scanStrings(H)

	// fmt.Println(inputs)
	res := createMap(inputs, H, W)
	// fmt.Println(res)

	for _, s := range res {
		fmt.Println(strings.Join(s, ""))
	}

}

func createMap(inputs [][]string, H int, W int) (res [][]string) {
	for i := 0; i < H; i++ {
		line := make([]string, 0)
		for j := 0; j < W; j++ {
			if inputs[i][j] == bomb {
				line = append(line, bomb)
				continue
			}
			point := 0
			// fmt.Println(i, j)
			if i-1 >= 0 {
				if j-1 >= 0 {
					if inputs[i-1][j-1] == bomb {
						// fmt.Println("a")
						point++
					}
				}
				if inputs[i-1][j] == bomb {
					// fmt.Println("b")
					point++
				}
				if j+1 <= W-1 {
					if inputs[i-1][j+1] == bomb {
						// fmt.Println("c")
						point++
					}
				}
			}
			if j-1 >= 0 {
				if inputs[i][j-1] == bomb {
					// fmt.Println("d")
					point++
				}
			}
			if j+1 <= W-1 {
				if inputs[i][j+1] == bomb {
					// fmt.Println("e")
					point++
				}
			}
			if i+1 <= H-1 {
				if j-1 >= 0 {
					if inputs[i+1][j-1] == bomb {
						// fmt.Println("f")
						point++
					}
				}
				if inputs[i+1][j] == bomb {
					// fmt.Println("g")
					point++
				}
				if j+1 <= W-1 {
					if inputs[i+1][j+1] == bomb {
						// fmt.Println("h")
						point++
					}
				}
			}
			// fmt.Println(point)
			line = append(line, strconv.Itoa(point))
		}

		// fmt.Println(line)
		res = append(res, line)
	}
	return
}

func scanStrings(len int) (inputs [][]string) {
	var str string
	for i := 0; i < len; i++ {
		fmt.Scanf("%s", &str)
		line := make([]string, 0)
		for _, s := range str {
			line = append(line, string(s))
		}

		inputs = append(inputs, line)
	}
	return
}
