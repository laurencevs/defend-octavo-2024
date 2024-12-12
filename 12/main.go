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

	grid := utils.GridFromBytes(mapData)
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

type region struct {
	plots []utils.Coord
	edges map[edgeKey][]int
	name  byte
}

type edgeKey struct {
	direction utils.Coord
	edgeIndex int // X index for vertical edge, Y index for horiziontal
}

func findRegions(grid utils.Grid) []region {
	seen := make(utils.Set[utils.Coord], grid.Length*grid.Width)
	coordStack := utils.NewStack[utils.Coord]()
	var regions []region
	for y, row := range grid.Values {
		for x, c := range row {
			coord := utils.Coord{Y: y, X: x}
			if seen.Contains(coord) {
				continue
			}
			currentRegion := region{[]utils.Coord{coord}, make(map[edgeKey][]int), c}
			coordStack.Add(coord)
			seen.Add(coord)
			for !coordStack.IsEmpty() {
				p := coordStack.Pop()
				for _, d := range utils.Directions {
					if adj := utils.AddCoords(p, d); !seen.Contains(adj) && grid.GetEntry(adj) == c {
						coordStack.Add(adj)
						seen.Add(adj)
						currentRegion.plots = append(currentRegion.plots, adj)
					}
				}
			}
			regions = append(regions, currentRegion)
		}
	}
	return regions
}

func computeStats(grid utils.Grid, regions []region) ([]int, []int, []int) {
	areas := make([]int, len(regions))
	perimeters := make([]int, len(regions))
	sideCounts := make([]int, len(regions))
	for i, region := range regions {
		areas[i] = len(region.plots)
		for _, plot := range region.plots {
			for _, d := range utils.Directions {
				if adj := utils.AddCoords(plot, d); grid.GetEntry(adj) != region.name {
					constantCoord := plot.X*utils.Abs(d.Y) + plot.Y*utils.Abs(d.X)
					edgeIndex := plot.X*utils.Abs(d.X) + plot.Y*utils.Abs(d.Y) +
						utils.Max(d.X+d.Y, 0) // for positive diff, index is plot index plus one
					key := edgeKey{direction: d, edgeIndex: edgeIndex}
					region.edges[key] = append(region.edges[key], constantCoord)
				}
			}
		}
		for _, edges := range region.edges {
			perimeters[i] += len(edges)
			edgeSet := utils.SetFromValues(edges...)
			var prevEdge bool
			for y := utils.Min(edges...); y <= utils.Max(edges...)+1; y++ {
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
