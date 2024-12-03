package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var mulPattern = regexp.MustCompile(`mul\(\d+,\d+\)`)
var mulDoPattern = regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don\'t\(\))`)

func scanString(s string) (int, error) {
	muls := mulDoPattern.FindAllString(s, -1)
	total := 0
	shouldMul := true
	for _, mul := range muls {
		if mul == "do()" {
			shouldMul = true
			continue
		} else if mul == "don't()" {
			shouldMul = false
			continue
		}
		if !shouldMul {
			continue
		}
		aStr, bStr, _ := strings.Cut(strings.Trim(mul, "mul()"), ",")
		a, err := strconv.Atoi(aStr)
		if err != nil {
			return 0, err
		}
		b, err := strconv.Atoi(bStr)
		if err != nil {
			return 0, err
		}
		total += a * b
	}
	return total, nil
}

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	output, err := scanString(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
