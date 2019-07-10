package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	alice := make([]int, 0, 100)
	bob := make([]int, 0, 100)

	cards := inputs[1:]
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	turn := 1
	for i := 0; i < len(cards); i++ {
		if turn%2 == 1 {
			alice = append(alice, cards[i])
		} else {
			bob = append(bob, cards[i])
		}
		turn++
	}

	aliceSum := 0
	for _, i := range alice {
		aliceSum = aliceSum + i
	}
	bobSum := 0
	for _, i := range bob {
		bobSum = bobSum + i
	}

	fmt.Println(aliceSum - bobSum)
}
