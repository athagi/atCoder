package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var s string
	list := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		list[i] = s
	}

	m := make(map[string]struct{})

    newList := make([]string, 0)

    for _, element := range list {
        if _, ok := m[element]; !ok {
            m[element] = struct{}{}
            newList = append(newList, element)
        }
	}
	
	fmt.Println(len(newList))
}
