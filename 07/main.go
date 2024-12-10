package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	candidateEquations, err := parseData(string(data))
	if err != nil {
		panic(err)
	}

	total1 := 0
	total2 := 0
	for _, eq := range candidateEquations {
		if canBeFixed(eq) {
			total1 += eq.target
		} else if canBeFixedP2(eq) {
			total2 += eq.target
		}
	}
	fmt.Println(total1)
	fmt.Println(total1 + total2)
}

func canBeFixed(eq candidateEquation) bool {
	for i := 0; i < 1<<(len(eq.operands)-1); i++ {
		v := eq.operands[0]
		for j, w := range eq.operands[1:] {
			if (i>>j)%2 == 0 {
				v += w
			} else {
				v *= w
			}
		}
		if v == eq.target {
			return true
		}
	}
	return false
}

func canBeFixedP2(eq candidateEquation) bool {
	for i := 0; i < 1<<(2*(len(eq.operands)-1)); i++ {
		v := eq.operands[0]
		for j, w := range eq.operands[1:] {
			switch (i >> (2 * j)) % 4 {
			case 0, 1:
				v += w
			case 2:
				v *= w
			case 3:
				v *= concatenationFactor(w)
				v += w
			}
		}
		if v == eq.target {
			return true
		}
	}
	return false
}

func concatenationFactor(x int) int {
	l := 10
	for x >= 10 {
		x /= 10
		l *= 10
	}
	return l
}

func parseData(data string) ([]candidateEquation, error) {
	lines := strings.Split(data, "\n")
	var equations []candidateEquation
	for _, line := range lines {
		targetStr, operandsStr, _ := strings.Cut(line, ": ")
		target, err := strconv.Atoi(targetStr)
		if err != nil {
			return nil, err
		}
		var operands []int
		for _, operandStr := range strings.Split(operandsStr, " ") {
			operand, err := strconv.Atoi(operandStr)
			if err != nil {
				return nil, err
			}
			operands = append(operands, operand)
		}
		equations = append(equations, candidateEquation{target, operands})
	}
	return equations, nil
}

type candidateEquation struct {
	target   int
	operands []int
}
