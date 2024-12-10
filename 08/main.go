package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	y, x int
}

func findAntinodes(data string, multiple bool) map[coord]struct{} {
	lines := strings.Split(data, "\n")
	antennaLocations := make(map[rune][]coord)
	height, width := len(lines), len(lines[0])
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				antennaLocations[char] = append(antennaLocations[char], coord{y, x})
			}
		}
	}
	antinodeLocations := make(map[coord]struct{})
	for _, antennas := range antennaLocations {
		for _, antenna1 := range antennas {
			for _, antenna2 := range antennas {
				if !multiple && antenna1 == antenna2 {
					continue
				}
				diff := coord{antenna2.y - antenna1.y, antenna2.x - antenna1.x}
				antinode := coord{antenna2.y + diff.y, antenna2.x + diff.x}
				for 0 <= antinode.y && antinode.y < height && 0 <= antinode.x && antinode.x < width {
					antinodeLocations[antinode] = struct{}{}
					if !multiple || antenna1 == antenna2 {
						break
					}
					antinode = coord{antinode.y + diff.y, antinode.x + diff.x}
				}
			}
		}
	}
	return antinodeLocations
}

func main() {
	mapData, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	antinodeLocations := findAntinodes(string(mapData), false)
	fmt.Println(len(antinodeLocations))
	antinodeLocations = findAntinodes(string(mapData), true)
	fmt.Println(len(antinodeLocations))
}
