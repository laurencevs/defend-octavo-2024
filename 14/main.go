package main

import (
	"aoc24/utils"
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
	g, err := parseData(string(data), os.Args[3], os.Args[4])
	if err != nil {
		panic(err)
	}

	g.moveRobots(100)
	fmt.Println(g.computeSafetyFactor())

	g.moveRobots(7038)
	err = os.WriteFile(os.Args[2], g.printRobots(), 0666)
	if err != nil {
		panic(err)
	}
}

type grid struct {
	h, w   int
	robots []robot
}

type robot struct {
	px, py int
	vx, vy int
}

func (g grid) printRobots() []byte {
	robotPositions := make(map[utils.Coord]int, len(g.robots))
	for _, r := range g.robots {
		robotPositions[utils.Coord{Y: r.py, X: r.px}]++
	}
	var output []byte
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			c := robotPositions[utils.Coord{Y: y, X: x}]
			var b byte
			switch {
			case c == 0:
				b = '.'
			case c >= 10:
				b = 'X'
			default:
				b = '0' + byte(c)
			}
			output = append(output, b)
		}
		output = append(output, '\n')
	}
	return output
}

func (g grid) computeSafetyFactor() int {
	var c [4]int
	for _, r := range g.robots {
		if r.px == g.w/2 || r.py == g.h/2 {
			continue
		}
		i := 0
		if r.px > g.w/2 {
			i += 2
		}
		if r.py > g.h/2 {
			i += 1
		}
		c[i]++
	}
	return c[0] * c[1] * c[2] * c[3]
}

func (g *grid) moveRobots(t int) {
	for i, r := range g.robots {
		g.robots[i].px += t * r.vx
		g.robots[i].px = (g.robots[i].px%g.w + g.w) % g.w
		g.robots[i].py += t * r.vy
		g.robots[i].py = (g.robots[i].py%g.h + g.h) % g.h
	}
}

func parseData(data, gridWidth, gridHeight string) (*grid, error) {
	var g grid
	var r robot
	var err error
	for _, line := range strings.Split(data, "\n") {
		_, err = fmt.Fscanf(strings.NewReader(line), "p=%d,%d v=%d,%d", &r.px, &r.py, &r.vx, &r.vy)
		if err != nil {
			return nil, err
		}
		g.robots = append(g.robots, r)
	}
	g.h, err = strconv.Atoi(gridHeight)
	if err != nil {
		return nil, err
	}
	g.w, err = strconv.Atoi(gridWidth)
	if err != nil {
		return nil, err
	}
	return &g, nil
}
