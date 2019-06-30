package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	inputs := make([]string, 0)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.TrimSpace(word)
		inputs = append(inputs, word)
	}

	var nums []int
	for _, s := range inputs[1:len(inputs)] {
		i, _ := strconv.Atoi(s)
		nums = append(nums, i)
	}
	var counter int
	flag := false
	for {
		var buf_nums []int
		for _, i := range nums {
			if i%2 == 0 {
				buf_nums = append(buf_nums, i/2)
			} else {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		nums = buf_nums
		counter++
	}

	fmt.Println(counter)
}
