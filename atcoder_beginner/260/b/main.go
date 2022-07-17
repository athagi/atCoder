package main

import (
	"fmt"
	"sort"
)

type person struct {
	id      int
	math    int
	english int
	sum     int
	passed  bool
}

func main() {
	var N, X, Y, Z, A, B int
	fmt.Scan(&N, &X, &Y, &Z)

	persons := make([]person, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		persons[i].id = i + 1
		persons[i].math = A
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&B)
		persons[i].english = B
		persons[i].sum = persons[i].math + B
		persons[i].passed = false
	}

	sort.SliceStable(persons, func(i, j int) bool { return persons[i].math > persons[j].math })
	for i := 0; i < X; i++ {
		persons[i].passed = true
	}

	sort.SliceStable(persons, func(i, j int) bool { return persons[i].id < persons[j].id })
	sort.SliceStable(persons, func(i, j int) bool { return persons[i].english > persons[j].english })
	for i := 0; i < N; i++ {
		if Y == 0 {
			break
		}
		if !persons[i].passed {
			persons[i].passed = true
			Y = Y - 1
		}
	}
	sort.SliceStable(persons, func(i, j int) bool { return persons[i].id < persons[j].id })
	sort.SliceStable(persons, func(i, j int) bool { return persons[i].sum > persons[j].sum })
	for i := 0; i < N; i++ {
		if Z == 0 {
			break
		}
		if !persons[i].passed {
			persons[i].passed = true
			Z = Z - 1
		}
	}

	sort.SliceStable(persons, func(i, j int) bool { return persons[i].id < persons[j].id })
	for i := 0; i < N; i++ {
		if persons[i].passed {

			fmt.Println(persons[i].id)
		}
	}
}
