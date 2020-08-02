package main

import (
    "fmt"

)



func main() {
	var N int
	var S, s1, s2 string
	fmt.Scan(&N, &S)
	i := 0
	j := N-1
	res := 0
	for {
		for {
			s1 = S[i: i+1]
			if (s1 == "W") {
				break
			}
			i++
			if (i > N-1){
				break
			}
		}
		if (i > N-1){
			break
		}

		
		for {
			s2 = S[j: j+1]
			if (s2 == "R") {
				break
			}
			j--
			if (j < 0) {
				break
			}
		}

		if i > j || j < 0{
			break
		} 



		res++
		i++
		j--
		
		
	}
	fmt.Println(res)

}
