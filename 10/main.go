package main

import (
	"aoc24/utils"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	coordsByHeight, topographicMap := parseData(data)

	total, uniqueTotal := countTrails(coordsByHeight, topographicMap)
	fmt.Println(total)
	fmt.Println(uniqueTotal)
}

type coord struct {
	y, x int
}

func parseData(data []byte) ([10][]coord, map[coord]int) {
	y, x := 0, 0
	coordsByHeight := [10][]coord{}
	topographicMap := make(map[coord]int)
	for _, c := range data {
		if c == '\n' {
			y += 1
			x = 0
			continue
		}
		topographicMap[coord{y, x}] = int(c - '0')
		coordsByHeight[c-'0'] = append(coordsByHeight[c-'0'], coord{y, x})
		x += 1
	}
	return coordsByHeight, topographicMap
}

var adj = []coord{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func countTrails(coordsByHeight [10][]coord, topographicMap map[coord]int) (int, int) {
	reachablePeaks := make(map[coord]utils.Set[coord], len(topographicMap))
	uniqueTrailCounts := make(map[coord]int, len(topographicMap))

	for _, c := range coordsByHeight[9] {
		reachablePeaks[c] = utils.Set[coord]{c: {}}
		uniqueTrailCounts[c] = 1
	}
	for h := 8; h >= 0; h-- {
		for _, c := range coordsByHeight[h] {
			reachablePeaks[c] = make(utils.Set[coord])
			for _, d := range adj {
				neighbour := coord{c.y + d.y, c.x + d.x}
				if topographicMap[neighbour] == h+1 {
					reachablePeaks[c].UnionUpdate(reachablePeaks[neighbour])
					uniqueTrailCounts[c] += uniqueTrailCounts[neighbour]
				}
			}
		}
	}

	var total, uniqueTotal int
	for _, c := range coordsByHeight[0] {
		total += len(reachablePeaks[c])
		uniqueTotal += uniqueTrailCounts[c]
	}
	return total, uniqueTotal
}
