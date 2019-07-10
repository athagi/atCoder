package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	if s[:1] != "d" && s[:1] != "e" {
		fmt.Println("NO")
		os.Exit(0)
	}

	for {
		var err error
		s, err = removeMatchedStr(s)
		if err != nil {
			fmt.Println("NO")
			os.Exit(0)
		}
		if len(s) == 0 {
			break
		}
	}
	fmt.Println("YES")
}

func removeMatchedStr(s string) (string, error) {
	// dreamRegex := regexp.MustCompile(`\w.dream$`)
	// dreamerRegex := regexp.MustCompile(`\w.dreamer$`)
	// eraseRegex := regexp.MustCompile(`\w.erase$`)
	// eraserRegex := regexp.MustCompile(`\w.eraser$`)
	if s == "dreamer" || s == "dream" || s == "eraser" || s == "erase" {
		return "", nil
	}

	// if dreamerRegex.MatchString(s) {
	if strings.HasSuffix(s, "dreamer") {
		return strings.TrimSuffix(s, "dreamer"), nil
	}

	// if dreamRegex.MatchString(s) {

	if strings.HasSuffix(s, "dream") {
		return strings.TrimSuffix(s, "dream"), nil
	}

	// if eraserRegex.MatchString(s) {

	if strings.HasSuffix(s, "eraser") {
		return strings.TrimSuffix(s, "eraser"), nil
	}

	// if eraseRegex.MatchString(s) {

	if strings.HasSuffix(s, "erase") {
		return strings.TrimSuffix(s, "erase"), nil
	}

	return "", errors.New("string is not found")
}
