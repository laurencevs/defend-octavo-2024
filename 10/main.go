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
	fmt.Println(countTrails(coordsByHeight, topographicMap))
	fmt.Println(countDistinctTrails(coordsByHeight, topographicMap))
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

func countDistinctTrails(coordsByHeight [10][]coord, topographicMap map[coord]int) int {
	trailCountMap := make(map[coord]int, len(topographicMap))
	for _, c := range coordsByHeight[9] {
		trailCountMap[c] = 1
	}
	for h := 8; h >= 0; h-- {
		for _, c := range coordsByHeight[h] {
			if topographicMap[coord{c.y + 1, c.x}] == h+1 {
				trailCountMap[c] += trailCountMap[coord{c.y + 1, c.x}]
			}
			if topographicMap[coord{c.y - 1, c.x}] == h+1 {
				trailCountMap[c] += trailCountMap[coord{c.y - 1, c.x}]
			}
			if topographicMap[coord{c.y, c.x + 1}] == h+1 {
				trailCountMap[c] += trailCountMap[coord{c.y, c.x + 1}]
			}
			if topographicMap[coord{c.y, c.x - 1}] == h+1 {
				trailCountMap[c] += trailCountMap[coord{c.y, c.x - 1}]
			}
		}
	}
	total := 0
	for _, c := range coordsByHeight[0] {
		total += trailCountMap[c]
	}
	return total
}

func countTrails(coordsByHeight [10][]coord, topographicMap map[coord]int) int {
	reachablePeaks := make(map[coord]utils.Set[coord], len(topographicMap))
	for _, c := range coordsByHeight[9] {
		reachablePeaks[c] = utils.Set[coord]{c: {}}
	}
	for h := 8; h >= 0; h-- {
		for _, c := range coordsByHeight[h] {
			reachablePeaks[c] = make(utils.Set[coord])
			if topographicMap[coord{c.y + 1, c.x}] == h+1 {
				reachablePeaks[c].UnionUpdate(reachablePeaks[coord{c.y + 1, c.x}])
			}
			if topographicMap[coord{c.y - 1, c.x}] == h+1 {
				reachablePeaks[c].UnionUpdate(reachablePeaks[coord{c.y - 1, c.x}])
			}
			if topographicMap[coord{c.y, c.x + 1}] == h+1 {
				reachablePeaks[c].UnionUpdate(reachablePeaks[coord{c.y, c.x + 1}])
			}
			if topographicMap[coord{c.y, c.x - 1}] == h+1 {
				reachablePeaks[c].UnionUpdate(reachablePeaks[coord{c.y, c.x - 1}])
			}
		}
	}
	total := 0
	for _, c := range coordsByHeight[0] {
		total += len(reachablePeaks[c])
	}
	return total
}
