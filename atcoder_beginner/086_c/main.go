package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	t int
	x int
	y int
}

type location struct {
	x int
	y int
}

func main() {

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	scanner.Split(bufio.ScanWords)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	inputs := make([]int, 0)
	for scanner.Scan() {
		word, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, word)
	}

	targets := inputs[1:]
	points := make([]point, 0)
	for i := 0; i < len(targets)/3; {
		points = append(points, point{t: targets[i], x: targets[i+1], y: targets[i+2]})
		i += 3
	}

	currentLocation := location{0, 0}
	currentTime := 0

	for _, p := range points {
		distance := checkDistance(currentLocation, p)
		remain := p.t - currentTime
		if distance <= remain && (remain-distance)%2 == 0 {
			currentTime = p.t
			currentLocation.x = p.x
			currentLocation.y = p.y
		} else {
			fmt.Println("No")
			os.Exit(0)
		}
	}
	fmt.Println("Yes")

}

func checkDistance(l location, p point) int {

	x := math.Abs(float64(p.x - l.x))
	y := math.Abs(float64(p.y - l.y))
	return int(x + y)
}
