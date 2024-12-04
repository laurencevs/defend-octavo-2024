package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func readLists(filename string) ([]int, []int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	lines := strings.Split(string(data), "\n")
	var list1, list2 []int
	for _, line := range lines {
		nos := strings.Split(line, "   ")
		n1, err := strconv.Atoi(nos[0])
		if err != nil {
			return nil, nil, err
		}
		n2, err := strconv.Atoi(nos[1])
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	return list1, list2, nil
}

func run() error {
	dataFilename := os.Args[1]
	list1, list2, err := readLists(dataFilename)
	if err != nil {
		return err
	}
	sort.Ints(list1)
	sort.Ints(list2)
	var totalDiff int
	for i := range list1 {
		totalDiff += diff(list1[i], list2[i])
	}
	fmt.Println(totalDiff)

	list2Counts := make(map[int]int, len(list2))
	for _, n := range list2 {
		list2Counts[n]++
	}
	similarity := 0
	for _, n := range list1 {
		similarity += n * list2Counts[n]
	}
	fmt.Println(similarity)
	return nil
}

func diff(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}
