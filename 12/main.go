package main

import (
	"aoc24/utils"
	"fmt"
	"os"
)

func main() {
	mapData, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	grid := parseMap(mapData)
	regions := findRegions(grid)
	areas, perimeters, sideCounts := computeStats(grid, regions)

	var cost1, cost2 int
	for i, a := range areas {
		cost1 += a * perimeters[i]
		cost2 += a * sideCounts[i]
	}
	fmt.Println(cost1)
	fmt.Println(cost2)
}

func parseMap(data []byte) [][]byte {
	grid := make([][]byte, 0)
	var row []byte
	for _, b := range data {
		if b == '\n' {
			grid = append(grid, row)
			row = nil
		} else {
			row = append(row, b)
		}
	}
	if row != nil {
		grid = append(grid, row)
	}
	return grid
}

type coord struct {
	y, x int
}

type region struct {
	plots           []coord
	horizontalEdges map[int][]int // y -> [x1, x2, ...], with x values inverted for north-facing edges
	verticalEdges   map[int][]int // x -> [y1, y2, ...], with x values inverted for west-facing edges
	char            byte
}

func findRegions(grid [][]byte) []region {
	inspected := make(utils.Set[coord], len(grid)*len(grid[0]))
	var inspectionStack *utils.Stack[coord]
	var regions []region
	var currentRegion region
	var p coord
	for y, row := range grid {
		for x, c := range row {
			if inspected.Contains(coord{y, x}) {
				continue
			}
			currentRegion = region{nil, make(map[int][]int), make(map[int][]int), c}
			inspectionStack = utils.NewStack[coord]()
			inspectionStack.Add(coord{y, x})
			for !inspectionStack.IsEmpty() {
				p = inspectionStack.Pop()
				if inspected.Contains(p) {
					continue
				}
				currentRegion.plots = append(currentRegion.plots, p)
				inspected.Add(p)
				if p.x > 0 && grid[p.y][p.x-1] == c {
					inspectionStack.Add(coord{p.y, p.x - 1})
				}
				if p.x < len(row)-1 && grid[p.y][p.x+1] == c {
					inspectionStack.Add(coord{p.y, p.x + 1})
				}
				if p.y > 0 && grid[p.y-1][p.x] == c {
					inspectionStack.Add(coord{p.y - 1, p.x})
				}
				if p.y < len(grid)-1 && grid[p.y+1][p.x] == c {
					inspectionStack.Add(coord{p.y + 1, p.x})
				}
			}
			regions = append(regions, currentRegion)
		}
	}
	return regions
}

func computeStats(grid [][]byte, regions []region) ([]int, []int, []int) {
	areas := make([]int, len(regions))
	perimeters := make([]int, len(regions))
	sideCounts := make([]int, len(regions))
	for i, region := range regions {
		areas[i] = len(region.plots)
		for _, plot := range region.plots {
			if plot.x == 0 || grid[plot.y][plot.x-1] != region.char {
				region.verticalEdges[plot.x] = append(region.verticalEdges[plot.x], -plot.y-1)
				perimeters[i]++
			}
			if plot.x == len(grid[0])-1 || grid[plot.y][plot.x+1] != region.char {
				region.verticalEdges[plot.x+1] = append(region.verticalEdges[plot.x+1], plot.y)
				perimeters[i]++
			}
			if plot.y == 0 || grid[plot.y-1][plot.x] != region.char {
				region.horizontalEdges[plot.y] = append(region.horizontalEdges[plot.y], -plot.x-1)
				perimeters[i]++
			}
			if plot.y == len(grid)-1 || grid[plot.y+1][plot.x] != region.char {
				region.horizontalEdges[plot.y+1] = append(region.horizontalEdges[plot.y+1], plot.x)
				perimeters[i]++
			}
		}
		for _, edges := range region.horizontalEdges {
			edgeSet := utils.SetFromValues(edges...)
			var prevEdge bool
			for x := -len(grid[0]); x <= len(grid[0])+1; x++ {
				hasEdge := edgeSet.Contains(x)
				if prevEdge && !hasEdge {
					sideCounts[i]++
				}
				prevEdge = hasEdge
			}
		}
		for _, edges := range region.verticalEdges {
			edgeSet := utils.SetFromValues(edges...)
			var prevEdge bool
			for y := -len(grid); y <= len(grid)+1; y++ {
				hasEdge := edgeSet.Contains(y)
				if prevEdge && !hasEdge {
					sideCounts[i]++
				}
				prevEdge = hasEdge
			}
		}
	}
	return areas, perimeters, sideCounts
}
