package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	inputs := scanInputs(N)
	sort.Stable(ByName(inputs))
	city := make([]Result, 0)
	city = append(city, inputs[0])
	for i := 1; i < N; i++ {
		if city[len(city)-1].City != inputs[i].City {
			sort.Sort(sort.Reverse(ByPoint(city)))
			for _, v := range city {
				fmt.Println(v.Num)
			}
			city = make([]Result, 0)
		}
		city = append(city, inputs[i])
	}
	sort.Sort(sort.Reverse(ByPoint(city)))
	for _, v := range city {
		fmt.Println(v.Num)
	}
	city = make([]Result, 0)
}

type ByPoint []Result

func (a ByPoint) Len() int           { return len(a) }
func (a ByPoint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoint) Less(i, j int) bool { return a[i].Point < a[j].Point }

type ByName []Result

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].City < b[j].City }

type Result struct {
	City  string
	Point int
	Num   int
}

func scanInputs(len int) (results []Result) {
	var city string
	var point int
	for i := 0; i < len; i++ {
		fmt.Scan(&city, &point)
		results = append(results, Result{city, point, i + 1})
	}
	return
}
