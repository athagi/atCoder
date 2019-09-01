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
	res := createMap(inputs, H, W)

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
			if i-1 >= 0 {
				if j-1 >= 0 {
					if inputs[i-1][j-1] == bomb {
						point++
					}
				}
				if inputs[i-1][j] == bomb {
					point++
				}
				if j+1 <= W-1 {
					if inputs[i-1][j+1] == bomb {
						point++
					}
				}
			}
			if j-1 >= 0 {
				if inputs[i][j-1] == bomb {
					point++
				}
			}
			if j+1 <= W-1 {
				if inputs[i][j+1] == bomb {
					point++
				}
			}
			if i+1 <= H-1 {
				if j-1 >= 0 {
					if inputs[i+1][j-1] == bomb {
						point++
					}
				}
				if inputs[i+1][j] == bomb {
					point++
				}
				if j+1 <= W-1 {
					if inputs[i+1][j+1] == bomb {
						point++
					}
				}
			}
			line = append(line, strconv.Itoa(point))
		}

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
