package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const offsetP2 = 10000000000000

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	problems, err := parseData(string(data))
	if err != nil {
		panic(err)
	}

	var cost int
	for _, p := range problems {
		cost += computeCost(p)
	}
	fmt.Println(cost)

	var cost2 int
	for _, p := range problems {
		p.px, p.py = p.px+offsetP2, p.py+offsetP2
		cost2 += computeCost(p)
	}
	fmt.Println(cost2)
}

type problem struct {
	ax, ay int
	bx, by int
	px, py int
}

func computeCost(p problem) int {
	aPushes, bPushes := invMul(p.ax, p.ay, p.bx, p.by, p.px, p.py)
	if aPushes*p.ax+bPushes*p.bx == p.px && aPushes*p.ay+bPushes*p.by == p.py {
		return 3*aPushes + bPushes
	}
	return 0
}

/*
Compute

	[ax bx]-1 [px]
	[ay by]   [py]

with integer (floor) division.
*/
func invMul(ax, ay, bx, by, px, py int) (int, int) {
	d := ax*by - ay*bx
	if d == 0 {
		return 0, 0
	}
	return (by*px - bx*py) / d, (ax*py - ay*px) / d
}

func parseData(data string) ([]problem, error) {
	var problems []problem
	var currentProblem problem
	for i, line := range strings.Split(data, "\n") {
		switch i % 4 {
		case 0, 1:
			xs, ys, _ := strings.Cut(line[12:], ", Y+")
			x, err := strconv.Atoi(xs)
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(ys)
			if err != nil {
				return nil, err
			}
			if i%4 == 0 {
				currentProblem.ax, currentProblem.ay = x, y
			} else {
				currentProblem.bx, currentProblem.by = x, y
			}
		case 2:
			pxs, pys, _ := strings.Cut(line[9:], ", Y=")
			px, err := strconv.Atoi(pxs)
			if err != nil {
				return nil, err
			}
			py, err := strconv.Atoi(pys)
			if err != nil {
				return nil, err
			}
			currentProblem.px, currentProblem.py = px, py
			problems = append(problems, currentProblem)
		}
	}
	return problems, nil
}
